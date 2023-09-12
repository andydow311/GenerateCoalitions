# GenerateSetofAgents

I am learning Golang and rereating a number of game theory problems to help me. 

This project generates coalitions of agents with values to reflect the synergy of cooperation.

The coalition value function is a data dei=riven function which reads data containing the pairwaise synergy and externality from co-operating and not cooperating together.

The bufio and OS libraries are used to read the input files.

To run:

  **go run main.go agents.go coalitions.go agents.txt functionData.txt**

To run tests:

**go mod init unit_test.go
go mod tidy
go test**
