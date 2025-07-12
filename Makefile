.PHONY: all backend frontend dev-frontend preview-frontend clean help

all: backend frontend

backend:
	cd ./ && go build -o savageapp

frontend:
	cd ./frontend && npm install && npm run build

dev-frontend:
	cd ./frontend && npm install && npm run dev

preview-frontend:
	cd ./frontend && npm run preview -- --port 3000

docker-build:
	docker build -t savageapp:latest .

clean:
	rm -f savageapp
	cd ./frontend && rm -rf build .svelte-kit node_modules

help:
	@echo "Available targets"
	@echo "  backend         Build Go backend"
	@echo "  frontend        Build Svelte frontend"
	@echo "  dev-frontend    Run Svelte frontend in dev mode"
	@echo "  preview-frontend Preview Svelte frontend on port 3000"
	@echo "  clean           Remove build artifacts"