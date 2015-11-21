// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddUser", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tUsers struct {}
var Users tUsers


func (_ tUsers) Login(
		login string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "login", login)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Users.Login", args).Url
}

func (_ tUsers) Signup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.Signup", args).Url
}

func (_ tUsers) SaveUser(
		login string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "login", login)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Users.SaveUser", args).Url
}


type tVoting struct {}
var Voting tVoting


func (_ tVoting) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Voting.List", args).Url
}

func (_ tVoting) ListDay(
		dayId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "dayId", dayId)
	return revel.MainRouter.Reverse("Voting.ListDay", args).Url
}

func (_ tVoting) VoteSlot(
		slotId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "slotId", slotId)
	return revel.MainRouter.Reverse("Voting.VoteSlot", args).Url
}

func (_ tVoting) SendVote(
		vote int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "vote", vote)
	return revel.MainRouter.Reverse("Voting.SendVote", args).Url
}


