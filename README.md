# KubeEase
k8s 管理平台后端

## 1. 服务镜像打包

### 克隆代码

```shell
$ git clone https://github.com/evescn/KubeEase.git
$ cd KubeEase
$ git checkout master
```

### 打包镜像

> 方法1 打包 Docker 镜像

```shell
# 第一种打包 Docker 镜像
# 编译项目
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# 打包 Docker 镜像
$ docker build -t harbor.xxx.cn/devops/kube-ease:v1.1 -f Dockerfile .
$ docker push harbor.xxx.cn/devops/kube-ease:v1.1
```

> 方法2 打包 Docker 镜像

```shell
# 第二种打包 Docker 镜像
$ chmod a+x ./build.sh
$ ./build.sh 1 dev # 版本号信息 环境
```

## 2. 服务部署

### a | Docker 启动

```shell
$ docker run -d \
  --restart=always \
  --name kube-ease \
  -p 7000:7000 \
  harbor.xxx.cn/devops/kube-ease:v1.1
```

### b | Kubernetes 启动

```yaml
# k8s.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-ease
  namespace: devops
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-ease
  template:
    metadata:
      labels:
        app: kube-ease
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - kube-ease
              topologyKey: kubernetes.io/hostname
            weight: 100
      containers:
      - name: kube-ease
        image: harbor.xxx.cn/devops/kube-ease:v1.1
        imagePullPolicy: Always
        ports:
        - containerPort: 7000
        livenessProbe:
          failureThreshold: 3
          initialDelaySeconds: 10
          periodSeconds: 5
          successThreshold: 1
          tcpSocket:
            port: 9000
          timeoutSeconds: 10
        readinessProbe:
          failureThreshold: 3
          initialDelaySeconds: 10
          periodSeconds: 5
          successThreshold: 1
          tcpSocket:
            port: 9000
          timeoutSeconds: 10
        resources:
          limits:
            cpu: "2"
            memory: 1Gi
          requests:
            cpu: 100m
            memory: 100Mi
        volumeMounts:
        - mountPath: /app/log/
          name: kube-ease-data
        - mountPath: /etc/localtime
          name: timezone
      volumes:
      - name: kube-ease-data
        persistentVolumeClaim:
          claimName: kube-ease-pvc
      - hostPath:
          path: /usr/share/zoneinfo/Asia/Shanghai
          type: ""
        name: timezone
---
# service
apiVersion: v1
kind: Service
metadata:
  name: kube-ease
  namespace: devops
spec:
  ports:
  - name: http
    nodePort: 27000
    port: 7000
    protocol: TCP
    targetPort: 7000
  - name: ws
    nodePort: 28082
    port: 8082
    protocol: TCP
    targetPort: 8082
  selector:
    app: kube-ease
  type: NodePort
---
# jenkins.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: kube-ease-pvc
  namespace: devops
spec:
  storageClassName: nfs-client-storageclass
  accessModes:
    - ReadWriteMany
  resources:
    requests: 
      storage: 50Gi   
```

```shell
$ kubectl apply -f k8s.yaml
```

## 3. 服务访问

> 项目前后端分离，需要部署前端后才能访问
> [前端地址](https://github.com/evescn/KubeEase-UI)


