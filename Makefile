
NAME=1337Med

all:
	go build -o $(NAME) main.go 

fclean:
	rm -rf $(NAME)
re: fclean all 