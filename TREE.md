.
├── Ai-todo.md
├── app
│   ├── controllers
│   │   ├── admin_controller.go
│   │   ├── admin_controller_test.go
│   │   ├── csv_import_controller.go
│   │   ├── csv_import_controller_test.go
│   │   ├── export_controller.go
│   │   ├── export_controller_test.go
│   │   ├── hello.go
│   │   ├── main_test.go
│   │   ├── role_controller.go
│   │   ├── role_controller_test.go
│   │   ├── status_controller.go
│   │   └── task_controller.go
│   ├── models
│   │   ├── admin.go
│   │   ├── exam_answer.go
│   │   └── role.go
│   └── services
│       ├── admin_service.go
│       ├── admin_service_test.go
│       ├── csv_importer.go
│       ├── csv_importer_test.go
│       ├── export_answer_service.go
│       ├── export_answer_service_test.go
│       ├── main_test.go
│       ├── mock_user_pool.go
│       ├── mock_user_pool_test.go
│       ├── redis_cleaner.go
│       ├── redis_cleaner_test.go
│       ├── redis_importer.go
│       ├── redis_importer_test.go
│       ├── redis_writer.go
│       ├── redis_writer_test.go
│       ├── role_service.go
│       ├── role_service_test.go
│       └── task_runner.go
├── ARCHITECTURE.md
├── auth
│   ├── auth.go
│   ├── middleware.go
│   └── permission.go
├── config
│   └── config.go
├── docs
│   ├── data
│   │   ├── full_schema_and_data.sql
│   │   ├── task.data.sql
│   │   ├── task_log.data.sql
│   │   ├── task_log.mk
│   │   ├── task.mk
│   │   ├── tm_admin.data.sql
│   │   ├── tm_admin.mk
│   │   ├── tm_exam_answers.data.sql
│   │   ├── tm_exam_answers.mk
│   │   ├── tm_exam_papers.data.sql
│   │   ├── tm_exam_papers.mk
│   │   ├── tm_exam_template.data.sql
│   │   ├── tm_exam_template.mk
│   │   ├── tm_permission.data.sql
│   │   ├── tm_permission.mk
│   │   ├── tm_role.data.sql
│   │   ├── tm_role.mk
│   │   ├── tm_role_permission.data.sql
│   │   ├── tm_role_permission.mk
│   │   ├── tm_user.data.sql
│   │   ├── tm_user.mk
│   │   ├── ym_admin.data.sql
│   │   ├── ym_admin.mk
│   │   ├── ym_answer.data.sql
│   │   ├── ym_answer.mk
│   │   ├── ym_answer_old.data.sql
│   │   ├── ym_answer_old.mk
│   │   ├── ym_article_contens.data.sql
│   │   ├── ym_article_contens.mk
│   │   ├── ym_article.data.sql
│   │   ├── ym_article_limit.data.sql
│   │   ├── ym_article_limit.mk
│   │   ├── ym_article.mk
│   │   ├── ym_cate_contens.data.sql
│   │   ├── ym_cate_contens.mk
│   │   ├── ym_cate.data.sql
│   │   ├── ym_cate.mk
│   │   ├── ym_guest.data.sql
│   │   ├── ym_guest.mk
│   │   ├── ym_link.data.sql
│   │   ├── ym_link.mk
│   │   ├── ym_member.data.sql
│   │   ├── ym_member.mk
│   │   ├── ym_score_stat.data.sql
│   │   └── ym_score_stat.mk
│   ├── data_import_export.md
│   ├── docs.go
│   ├── GEN_CURD.md
│   ├── goapp_structure.md
│   ├── swagger.json
│   └── swagger.yaml
├── gin.log
├── go.mod
├── go.sum
├── handlers
│   ├── dbinfo.go
│   ├── exam_answer.go
│   ├── exam.go
│   ├── exam_paper.go
│   ├── exam_paper_redis.go
│   ├── exam_template.go
│   ├── hello.go
│   ├── source_check.go
│   ├── status.go
│   ├── upload.go
│   └── version.go
├── hello
├── hello.log
├── LICENSE.md
├── main.go
├── output.log
├── package.json
├── package-lock.json
├── README.md
├── routes
│   └── routes.go
├── run.sh
├── sql.sh
├── static
│   ├── exam_template.html
│   ├── exam_write.html
│   ├── exports
│   │   ├── score_0520081348_d5db.csv
│   │   └── score_0520081423_4159.csv
│   ├── hello.html
│   ├── index.html
│   └── uploads
│       ├── images
│       │   ├── exam
│       │   │   ├── a.jpg
│       │   │   ├── exam_0225465e-f34e-4720-8600-0b40fe0e10ea.jpg
│       │   │   ├── exam_051d104a-2982-41ff-841a-361521c82422.jpg
│       │   │   ├── exam_0915fe55-2288-467b-a743-513bfec7978f.jpg
│       │   │   ├── exam_0dc70d79-ba26-498f-aca2-a335e81fabf8.png
│       │   │   ├── exam_0ef1fe02-6a46-4a41-aaae-43531430ca6b.jpg
│       │   │   ├── exam_1f9402d3-026b-42eb-bb01-7fb00dae5114.jpg
│       │   │   ├── exam_2010e39a-6859-4af1-8361-e0105f3c752f.jpg
│       │   │   ├── exam_2313cafe-106e-4f03-b543-c8bce102b992.jpg
│       │   │   ├── exam_28a545f5-36c7-44bb-99ca-bb401d46c564.jpg
│       │   │   ├── exam_29d6018a-9753-488f-9c92-610553d70a23.jpg
│       │   │   ├── exam_2a016b37-4790-4d5d-ab5f-a30bd23315b7.jpg
│       │   │   ├── exam_30be087c-f4c7-4428-b0ee-548a63beb8b6.jpg
│       │   │   ├── exam_36e36d5f-919b-499d-a133-91b564a2f816.jpg
│       │   │   ├── exam_3a39a20c-0fdd-441a-b36b-680d13f3fcce.jpg
│       │   │   ├── exam_3bd95a45-5922-49dd-8616-a498e8772773.png
│       │   │   ├── exam_40458c85-6a0e-4129-80db-cbea77f0c4fc.jpg
│       │   │   ├── exam_40eeb196-6a12-405e-bc7f-c1c39fed9881.jpg
│       │   │   ├── exam_4ae237a3-43b7-43d6-a34e-9fe30e1a25e7.jpg
│       │   │   ├── exam_4c2cb46e-6feb-4834-a0be-59e91c19ceb7.jpg
│       │   │   ├── exam_4ccd2670-213f-4623-8383-525fc1f05e4c.jpg
│       │   │   ├── exam_5be5b228-56ca-422c-a137-21f3ae4fc529.jpg
│       │   │   ├── exam_6062ded1-8011-4173-9b07-0e17a3ce2ff0.jpg
│       │   │   ├── exam_60ee7446-9eeb-4fa2-a318-068a8530a4de.png
│       │   │   ├── exam_717a9368-aed7-49cc-9741-c9b332f1abc7.jpg
│       │   │   ├── exam_733d1b8a-a6d2-428c-b3fb-d4d1c758009a.jpg
│       │   │   ├── exam_77952e20-d1c8-409b-9b94-4c7fda645f85.jpg
│       │   │   ├── exam_814dac8a-8b25-4aec-bad8-03b9660b8b77.jpg
│       │   │   ├── exam_82d373d6-1041-4a5c-9855-5ca5f50e55a7.jpg
│       │   │   ├── exam_8c35a925-770c-4737-afbf-c6070bf30e29.jpg
│       │   │   ├── exam_900379c3-3b65-497e-9b31-b2a8cc9450e1.jpg
│       │   │   ├── exam_936d68e1-1ed9-4878-b60b-527610580ea4.jpg
│       │   │   ├── exam_93b33c34-b757-4735-88d2-478e61a4b40d.jpg
│       │   │   ├── exam_9c17d443-598f-4dd9-8c7d-ddd2e9abba6b.jpg
│       │   │   ├── exam_9ed09533-7158-4022-9144-7a2be9ef4923.jpg
│       │   │   ├── exam_a2dc9767-6628-468e-b5ca-3b81add65feb.jpg
│       │   │   ├── exam_a73580cf-5d0d-4065-bd12-d1da9cc5b8a9.jpg
│       │   │   ├── exam_a7897895-0f18-41c6-9d05-194bdfb94c4f.jpg
│       │   │   ├── exam_aae84063-5f24-49ca-a0e6-74f8b6f77fa1.png
│       │   │   ├── exam_abfba4d7-27f8-4894-b89c-e8d5f7b30a44.jpg
│       │   │   ├── exam_ac500aec-38bf-47cc-b04c-0e2500442c1c.png
│       │   │   ├── exam_ac6e4e97-3e82-4497-ba01-ef62dde485d3.png
│       │   │   ├── exam_bb78c594-6479-45ff-93fa-d9cd9c47ff56.png
│       │   │   ├── exam_bfeee036-b4b8-4176-a867-8f91fa3c0a56.jpg
│       │   │   ├── exam_c17a9f55-d9ac-422c-a9d6-58677aae2416.jpg
│       │   │   ├── exam_c534b1a9-f59a-4b4d-8265-4e553186cbca.jpg
│       │   │   ├── exam_dd035bf9-6c44-467c-90be-1bf723ead66e.jpg
│       │   │   ├── exam_e29e2780-8b7c-4bdf-ad17-70d91f3b38d0.jpg
│       │   │   ├── exam_e73b5ff4-32b2-4718-baea-c80cc80473fd.jpg
│       │   │   ├── exam_f0c54536-b9e0-4f7b-8227-75ad9f3f9479.jpg
│       │   │   ├── exam_f24afe9b-d1ae-4ab2-90a0-a2c8637e050e.png
│       │   │   ├── exam_f3fde2db-e9ad-4a6d-8c60-761e2dc50968.jpg
│       │   │   ├── exam_f54271a4-1278-477e-a140-7fc87748fe4e.jpg
│       │   │   ├── exam_f6960a42-b080-43b2-8c5a-9eacfbb985c8.jpg
│       │   │   └── index.html
│       │   └── index.html
│       ├── index.html
│       └── score_0519103002_30cd.csv
├── TREE.md
├── tree.sh
└── utils
    ├── db.go
    ├── db_gorm.go
    ├── db_sqlx.go
    ├── jwt.go
    ├── password.go
    ├── path.go
    ├── queue.go
    ├── redis.go
    ├── status.go
    └── time.go

16 directories, 191 files
