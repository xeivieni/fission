package main

import (
	"github.com/posener/complete"
	"github.com/fission/fission/controller/client"
	"os"
	"fmt"
)

func setupControllerClient() *client.Client {
	serverUrl := os.Getenv("FISSION_URL")
	return client.MakeClient(serverUrl)
}

var PredictFunctionList = complete.PredictFunc(func(complete.Args) []string {
	client := setupControllerClient()
	//fmt.Println("Setup controller client")

	fnList, err := client.FunctionList()
	if err != nil {
		fmt.Printf("get function list, err : %v", err)
		return nil
	}

	funcList := make([]string, 0)
	for _, fn := range fnList {
		funcList= append(funcList, fn.Metadata.Name)
	}

	return funcList
})

var PredictEnvList = complete.PredictFunc(func(complete.Args) []string {
	client := setupControllerClient()
	//fmt.Println("Setup controller client")

	envrList, err := client.EnvironmentList()
	if err != nil {
		fmt.Printf("error getting environment list, err : %v", err)
		return nil
	}

	envList := make([]string, 0)
	for _, env := range envrList {
		envList = append(envList, env.Metadata.Name)
	}

	return envList
})


func main() {

	// create a Command object, that represents the command we want
	// to complete.
	run := complete.Command{

		// Sub defines a list of sub commands of the program,
		// this is recursive, since every command is of type command also.
		Sub: complete.Commands{

			// add a build sub command
			"function": complete.Command{

				// Sub defines a list of sub commands of the program,
				// this is recursive, since every command is of type command also.
				Sub: complete.Commands{

					// add a create sub command
					"create": complete.Command{

						//minCpu, maxCpu, minMem, maxMem, minScale, maxScale, fnExecutorTypeFlag, targetcpu, fnCfgMapFlag, fnSecretFlag, fnSecretnsFlag, fnCfgMapnsFlag}
						// define flags of the create sub command
						Flags: complete.Flags{
							// build sub command has a flag '-cpus', which
							// expects number of cpus after it. in that case
							// anything could complete this flag.
							"--name": complete.PredictAnything,
							"--env": PredictEnvList,
							"--spec": complete.PredictAnything,
							"--code": complete.PredictFiles("*"),
							"--src": complete.PredictAnything,
							"--deploy": complete.PredictAnything,
							"--entrypoint": complete.PredictAnything,
							"--buildcmd": complete.PredictAnything,
							"--pkg": complete.PredictAnything,
							"--url": complete.PredictAnything,
							"--method": complete.PredictAnything,
						},
					},

					// add a create sub command
					"update": complete.Command {

						//minCpu, maxCpu, minMem, maxMem, minScale, maxScale, fnExecutorTypeFlag, targetcpu, fnCfgMapFlag, fnSecretFlag, fnSecretnsFlag, fnCfgMapnsFlag}
						// define flags of the create sub command
						Flags: complete.Flags{
							// build sub command has a flag '-cpus', which
							// expects number of cpus after it. in that case
							// anything could complete this flag.
							"--name": PredictFunctionList,
							"--env": PredictEnvList,
							"--spec": complete.PredictAnything,
							"--code": complete.PredictFiles("*"),
							"--src": complete.PredictAnything,
							"--deploy": complete.PredictAnything,
							"--entrypoint": complete.PredictAnything,
							"--buildcmd": complete.PredictAnything,
							"--pkg": complete.PredictAnything,
							"--url": complete.PredictAnything,
							"--method": complete.PredictAnything,
						},
					},

					"list" : complete.Command{},

					// add a create sub command
					"delete": complete.Command {

						//minCpu, maxCpu, minMem, maxMem, minScale, maxScale, fnExecutorTypeFlag, targetcpu, fnCfgMapFlag, fnSecretFlag, fnSecretnsFlag, fnCfgMapnsFlag}
						// define flags of the create sub command
						Flags: complete.Flags{
							// build sub command has a flag '-cpus', which
							// expects number of cpus after it. in that case
							// anything could complete this flag.
							"--name": PredictFunctionList,
						},
					},
				},
			},

			// Sub defines a list of sub commands of the program,
			// this is recursive, since every command is of type command also.

			// add a build sub command
			"env": complete.Command{

				// Sub defines a list of sub commands of the program,
				// this is recursive, since every command is of type command also.
				Sub: complete.Commands{

					// add a create sub command
					"create": complete.Command{

						//minCpu, maxCpu, minMem, maxMem, minScale, maxScale, fnExecutorTypeFlag, targetcpu, fnCfgMapFlag, fnSecretFlag, fnSecretnsFlag, fnCfgMapnsFlag}
						// define flags of the create sub command
						Flags: complete.Flags{
							// build sub command has a flag '-cpus', which
							// expects number of cpus after it. in that case
							// anything could complete this flag.
							"--name": complete.PredictAnything,
							"--image": complete.PredictAnything,
						},
					},

					"list" : complete.Command{},

					// add a create sub command
					"delete": complete.Command {

						//minCpu, maxCpu, minMem, maxMem, minScale, maxScale, fnExecutorTypeFlag, targetcpu, fnCfgMapFlag, fnSecretFlag, fnSecretnsFlag, fnCfgMapnsFlag}
						// define flags of the create sub command
						Flags: complete.Flags{
							// build sub command has a flag '-cpus', which
							// expects number of cpus after it. in that case
							// anything could complete this flag.
							"--name": PredictEnvList,
						},
					},
				},
			},
		},
	}


	// run the command completion, as part of the main() function.
	// this triggers the autocompletion when needed.
	// name must be exactly as the binary that we want to complete.
	complete.New("fission", run).Run()
}