package main

import (
	"fmt"
	"net/http"
	"tellMeWhen/query"
	"tellMeWhen/remind"
)

func main() {
	//var conStr = "root:root@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	//db, _ := gorm.Open(mysql.Open(conStr), &gorm.Config{})
	//// specify the output directory (default: "./query")
	//// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	//g := gen.NewGenerator(gen.Config{
	//	OutPath: "/dal/query",
	//	/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
	//	//if you want the nullable field generation property to be pointer type, set FieldNullable true
	//	/* FieldNullable: true,*/
	//	//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
	//	/* FieldCoverable: true,*/
	//	//if you want to generate index tags from database, set FieldWithIndexTag true
	//	/* FieldWithIndexTag: true,*/
	//	//if you want to generate type tags from database, set FieldWithTypeTag true
	//	/* FieldWithTypeTag: true,*/
	//	//if you need unit tests for query code, set WithUnitTest true
	//	/* WithUnitTest: true, */
	//})
	//
	//// reuse the database connection in Project or create a connection here
	//// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	//// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	//g.UseDB(db)
	//
	//// apply basic crud api on structs or table models which is specified by table name with function
	//// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	////g.ApplyBasic(model.Reminder{}, g.GenerateModel("reminder"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	//
	//// apply diy interfaces on structs or table models
	////g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))
	//
	//// execute the action of code generation
	//g.Execute()
	rq := query.GetRemindQuery()
	reminder := rq.GetAllReminder()
	fmt.Println(reminder)
	sender := remind.GetSender()
	for i := 0; i < len(reminder); i++ {
		r := reminder[i]
		sender.AddReminder(remind.GetReminderInterfaceByModel(r))
	}
	err := http.ListenAndServe(":8555", nil)
	if err != nil {
		fmt.Println("listen fail...")
	}
}
