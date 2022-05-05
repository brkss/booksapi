
NAME 	= booksapi

OBJ 	= cmd/main/main.go


all : $(NAME)

$(NAME):
	@echo "⚙️  Compiling..."
	go build -o $(NAME) $(OBJ)

run: $(NAME)
	@echo "🍭 Runing $(NAME) ..."
	./$(NAME)

clean:
	@echo "🧹 Cleaning..."
	rm -f $(NAME)

re: clean all

.PHONY: clean compile run all
	
