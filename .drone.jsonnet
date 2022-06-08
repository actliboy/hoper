// local mode(mode="app") = if mode == "app" then "app" else "node";
local tpldir = "./build/k8s/app/";
local codedir = "/mnt/new/code/hoper/";

local kubectl (deplocal,cmd) = if deplocal then{
      name: "deploy",
      image: "bitnami/kubectl",
      user: 0, //文档说是string类型，结果"root"不行 k8s runAsUser: 0
      volumes: [
        {
          name: "kube",
          path: "/root/.kube/"
        },
        {
          name: "minikube",
          path: "/root/.minikube/"
        }
      ],
      commands:cmd
} else {
      name: "deploy",
      image: "bitnami/kubectl",
      user: 0, //文档说是string类型，结果"root"不行 k8s runAsUser: 0
      environment:  {
             CA: {
                from_secret: "ca"
              },
              CACRT: {
                from_secret: "ca_crt"
              },
               CAKEY: {
                  from_secret: "ca_key"
               },
      },
      commands:[
            "chmod +x "+ tpldir+"account.sh && "+ tpldir+"account.sh",
      ]+cmd
};


local Pipeline(group, name, mode, protoc, workdir, sourceFile="", opts=[],deplocal=false,schedule="") = {
  local fullname = if name == "" then group else group + "-" + name,
  local tag = "${DRONE_TAG##"+fullname+"-v}",
  local datadir = if deplocal then "/mnt/new/data" else "/data",
  local dockerfilepath = tpldir+mode+"/Dockerfile",
  local deppath = tpldir+mode+"/deployment.yaml",
  kind: "pipeline",
  type: "kubernetes",
  name: fullname,
  metadata: {
    namespace: "default"
  },
  platform: {
    os: "linux",
    arch: "amd64"
  },
  trigger: {
    ref: [
      "refs/tags/"+fullname+"-v*"
      ]
  },
  volumes: [
    {
        name: "codedir",
        host: {
           path: codedir
        }
    },
    {
      name: "gopath",
      host: {
         path: datadir+"/deps/gopath/"
      }
    },
     {
       name: "dockersock",
       host: {
         path: "/var/run/"
       }
     },
     {
       name: "kube",
       host: {
           path: "/root/.kube/"
       }
     },
      {
        name: "minikube",
        host: {
          path: "/root/.minikube/"
        }
      }
  ],
  clone: {
   disable: true
  },
  steps: [
    {
      name: "clone",
      image: "alpine/git",
      volumes: [
          {
              name: "codedir",
              path: "/code/"
          }
      ],
      commands: [
        "git config --global http.proxy 'socks5://proxy.tools:1080'",
        "git config --global https.proxy 'socks5://proxy.tools:1080'",
        //"git clone ${DRONE_GIT_HTTP_URL} .",
        "cd /code",
        "git tag -l | xargs git tag -d",
        "git fetch --all && git reset --hard origin/master && git pull",
        "cd /drone/src/",
        "git clone /code .",
        "git checkout -b deploy $DRONE_COMMIT_REF",
        local buildfile = "/code/"+workdir+"/protobuf/build";
        if protoc then "if [ -f "+buildfile+" ]; then cp -r /code/"+workdir+"/protobuf  /drone/src/"+workdir+"; fi" else "echo",
        "sed -i 's/$${app}/"+fullname+"/g' "+dockerfilepath,
        local cmd = ["./"+fullname]+opts;
        "sed -i 's#$${cmd}#"+std.join(" ,",["\""+opt+"\"" for opt in cmd])+"#g' "+dockerfilepath,
        "sed -i 's/$${app}/"+fullname+"/g' "+deppath,
        "sed -i 's/$${group}/"+group+"/g' "+deppath,
        "sed -i 's#$${datadir}#"+datadir+"#g' "+deppath,
        "sed -i 's#$${image}#jyblsq/"+fullname+":"+tag+"#g' "+deppath,
        if  mode == "cronjob"  then "sed -i 's#$${schedule}#"+schedule+"#g' "+ deppath else "echo",
        "cp -r "+deppath+" /code/deploy/"+mode+"/"+fullname+"-"+tag+".yaml",
      ]
    },
    {
      name: "go build",
      image: if protoc then "jyblsq/golang:protoc" else "golang:1.18.1",
      volumes: [
        {
            name: "gopath",
            path: "/go/"
        }
      ],
      environment: {
         GOPROXY: "https://goproxy.io,https://goproxy.cn,direct"
      },
      commands: [
        "cd " + workdir,
        "go mod download",
        local buildfile = "/drone/src/"+workdir+"/protobuf/build";
         if protoc then "if [ ! -f "+buildfile+" ]; then go run ./protobuf; fi" else "echo",
        "go mod tidy",
        "go build -o  /drone/src/"+fullname+" "+sourceFile
      ]
    },
    {
      name: "docker build",
      image: "plugins/docker",
      privileged: true ,
      volumes: [
        {
            name: "dockersock",
            path: "/var/run/"
        }
      ],
      settings: {
        username: {
          from_secret: "docker_username"
        },
        password: {
          from_secret: "docker_password"
        },
       repo: "jyblsq/"+fullname,
       tags: tag,
       dockerfile: dockerfilepath,
       force_tag: true,
       auto_tag: false,
       daemon_off: true,
       purge: true,
       pull_image: false,
       dry_run: deplocal
      }
    },
    kubectl(deplocal,[
        if tag != "1.0.0" && (mode == "job" || mode == "cronjob")  then "kubectl --kubeconfig=/root/.kube/config delete -f "+ deppath else "echo",
       "kubectl --kubeconfig=/root/.kube/config apply -f "+deppath,
    ]),
    {
       name: "dingtalk",
       image: "lddsb/drone-dingtalk-message",
       settings: {
          token: {
            from_secret: "token"
          },
          type: "markdown",
          message_color: true,
          message_pic: true,
          sha_link: true
       }
    }
  ]
};

[
  Pipeline("timepill","","app",false,"tools/server","./timepill/cmd/record.go",["-t"]),
  Pipeline("hoper","","app",true,"server/go/mod"),
  Pipeline("timepill","rbyorderid","job",false,"tools/server","./timepill/cmd/recordby_orderid.go"),
  Pipeline("timepill","esload","cronjob",false,"tools/server","./timepill/cmd/search_es8.go",deplocal=true, schedule="0 23 * * *"),
]