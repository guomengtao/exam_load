gin-go-test/
├── app/
│   ├── controllers/
│   ├── models/
│   ├── services/
│   └── biz/
│
├── routes/
│   ├── routes.go
│   └── gen_routes.go
│
├── utils/
│   ├── gen/
│   │   ├── gen.go
│   │   ├── templates/          # 模板文件放这里
│   │   │   ├── model.tpl
│   │   │   ├── service.tpl
│   │   │   ├── biz.tpl
│   │   │   ├── controller.tpl
│   │   │   ├── route.tpl
│   │   │   └── ...
│   ├── genlib/
│   │   ├── gen_model.go
│   │   ├── gen_service.go
│   │   ├── gen_biz.go
│   │   ├── gen_controller.go
│   │   ├── gen_route.go
│   │   └── gen_utils.go
│
├── utils/generated/
│   ├── models/
│   ├── services/
│   ├── controllers/
│   └── routes/
│
├── main.go
├── go.mod
├── go.sum
└── README.md