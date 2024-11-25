
NAME=cppcreate

all:
	go build -o $(NAME) main.go 

fclean:
	rm -rf $(NAME)
re: fclean all 