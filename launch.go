package main

import (
	//    "io"
	//   "os"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	//    "github.com/docker/docker/pkg/stdcopy"
)

func main() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	//_, err = cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	//if err != nil {
	//    panic(err)
	//}
	var cont_id []string

	for i := 0; i < 50; i++ {

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: "alpine",
			Cmd:   []string{"echo", "hello world"},
		}, nil, nil, "")
		if err != nil {
			panic(err)
		}
		cont_id = append(cont_id, resp.ID)
		//fmt.Println(cont_id)
		//fmt.Println(resp)
		//fmt.Println(ctx)
	}

	var wg sync.WaitGroup
	wg.Add(len(cont_id)) // going to run number of threads=containers
	fmt.Println(len(cont_id))
	start := time.Now()
	for i := 0; i < 50; i++ {
		go func(i int) {
			defer wg.Done()

			if err := cli.ContainerStart(ctx, cont_id[i], types.ContainerStartOptions{}); err != nil {
				panic(err)
			}

			//statusCh, errCh := cli.ContainerWait(ctx, cont_id[i], container.WaitConditionNotRunning)
			//select {
			//case err := <-errCh:
			//	if err != nil {
			//		panic(err)
			//	}
			//case <-statusCh:
			//}

			//out, err := cli.ContainerLogs(ctx, cont_id[i], types.ContainerLogsOptions{ShowStdout: true})
			//if err != nil {
			//    panic(err)
			//}

			//stdcopy.StdCopy(os.Stdout, os.Stderr, out)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	//fmt.Println("Finished for loop")
}
