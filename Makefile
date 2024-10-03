DEFAULT_STEP := 0
GOPATH := $(shell go env GOPATH)
step:=$(or $(step), $(DEFAULT_STEP))
# Admins
# -------------------------------------------------------------------------
serve-admins:
	cd services/admin && \
	go run ./... serve
migrate-up-admins:
	cd services/admin && \
	go run ./... migrate --action=up --step=$(step)
migrate-down-admins:
	cd services/admin && \
	go run ./... migrate --action=down --step=$(step)
create-admin:
	cd services/admin && \
	go run ./... createAdmin --email=$(email) --password=$(password) --name=$(name)
test-admins:
	cd services/admin && \
	ENV="test" && go run ./... migrate && go test ./... -v
protoc-admins:
	cd services/admin && \
	protoc -I protos protos/v1/admin.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen
# Customers
# -------------------------------------------------------------------------
migrate-up-customers:
	cd services/customer && \
	go run ./... migrate --action=up  --step=$(step)
migrate-down-customers:
	cd services/customer && \
	go run ./... migrate --action=down  --step=$(step)
protoc-customers:
	cd services/customer && \
	protoc -I protos protos/v1/customer.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen
serve-customers:
	cd services/customer && \
	go run ./... serve
# Orders
# -------------------------------------------------------------------------
migrate-up-orders:
	cd services/order && \
	go run ./... migrate --action=up --step=$(step)
migrate-down-orders:
	cd services/order && \
	go run ./... migrate --action=down --step=$(step)
protoc-orders:
	cd services/order && \
	protoc -I protos protos/v1/customer.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen &&\
	protoc -I protos protos/v1/product.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen &&\
	protoc -I$(GOPATH)/src/github.com/googleapis/googleapis -I protos protos/v1/payment.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen &&\
	protoc -I protos protos/v1/order.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen
serve-orders:
	cd services/order && \
	go run ./... serve
# Products
# -------------------------------------------------------------------------
protoc-products:
	cd services/product && \
	protoc -I protos protos/v1/product.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen && \
	protoc -I protos protos/v1/admin.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen

serve-products:
	cd services/product && \
	go run ./... serve
# pyaments
# -------------------------------------------------------------------------
migrate-up-payments:
	cd services/payment && \
	go run ./... migrate --action=up --step=$(step)
migrate-down-payments:
	cd services/payment && \
	go run ./... migrate --action=down --step=$(step)
protoc-payments:
	cd services/payment && \
	protoc -I$(GOPATH)/src/github.com/googleapis/googleapis -I protos protos/v1/payment.proto --go-grpc_out=protos/v1/gen --go_out=protos/v1/gen
serve-payments:
	cd services/payment && \
	go run ./... serve
# API Gateway
# -------------------------------------------------------------------------
serve-api-gateway:
	cd services/api_gateway && \
	go run ./... serve
protoc-api-gateway:
	cd services/api_gateway && \
	protoc -I . \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	-I$(GOPATH)/src/github.com/googleapis/googleapis \
	--go_out ./protos/v1/gen/admin \
	--go-grpc_out ./protos/v1/gen/admin \
	--grpc-gateway_opt=grpc_api_configuration=protos/v1/admin-proto-conf.yaml \
	--grpc-gateway_out=./protos/v1/gen/admin \
	--grpc-gateway_opt=logtostderr=true \
	--openapiv2_out ./protos/v1/swagger \
	--openapiv2_opt=grpc_api_configuration=protos/v1/admin-proto-conf.yaml \
	protos/v1/admin.proto \
	&& 	protoc -I . \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	-I$(GOPATH)/src/github.com/googleapis/googleapis \
	--go_out ./protos/v1/gen/product \
	--go-grpc_out ./protos/v1/gen/product \
	--grpc-gateway_opt=grpc_api_configuration=protos/v1/product-proto-conf.yaml \
	--grpc-gateway_out=./protos/v1/gen/product \
	--grpc-gateway_opt=logtostderr=true \
	--openapiv2_out ./protos/v1/swagger \
	--openapiv2_opt=grpc_api_configuration=protos/v1/product-proto-conf.yaml \
	protos/v1/product.proto \
	&& 	protoc -I . \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	-I$(GOPATH)/src/github.com/googleapis/googleapis \
	--go_out ./protos/v1/gen/order \
	--go-grpc_out ./protos/v1/gen/order \
	--grpc-gateway_opt=grpc_api_configuration=protos/v1/order-proto-conf.yaml \
	--grpc-gateway_out=./protos/v1/gen/order \
	--grpc-gateway_opt=logtostderr=true \
	--openapiv2_out ./protos/v1/swagger \
	--openapiv2_opt=grpc_api_configuration=protos/v1/order-proto-conf.yaml \
	protos/v1/order.proto \
	&&	protoc -I . \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	-I$(GOPATH)/src/github.com/googleapis/googleapis \
	--go_out ./protos/v1/gen/customer \
	--go-grpc_out ./protos/v1/gen/customer \
	--grpc-gateway_opt=grpc_api_configuration=protos/v1/customer-proto-conf.yaml \
	--grpc-gateway_out=./protos/v1/gen/customer \
	--grpc-gateway_opt=logtostderr=true \
	--openapiv2_out ./protos/v1/swagger \
	--openapiv2_opt=grpc_api_configuration=protos/v1/customer-proto-conf.yaml \
	protos/v1/customer.proto \
	&&	protoc -I . \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	-I$(GOPATH)/src/github.com/googleapis/googleapis \
	--go_out ./protos/v1/gen/payment \
	--go-grpc_out ./protos/v1/gen/payment \
	--grpc-gateway_opt=grpc_api_configuration=protos/v1/payment-proto-conf.yaml \
	--grpc-gateway_out=./protos/v1/gen/payment \
	--grpc-gateway_opt=logtostderr=true \
	--openapiv2_out ./protos/v1/swagger \
	--openapiv2_opt=grpc_api_configuration=protos/v1/payment-proto-conf.yaml \
	protos/v1/payment.proto

# Minikube K8S (local environment)
# -------------------------------------------------------------------------
minikube-k8s:
	rm E-Commerce-0.1.0.tgz > /dev/null 2>&1 || true && \
	helm delete ecommerce > /dev/null 2>&1 || true && \
	docker build -t customer:$(version) -f deployments/docker/app-go.Dockerfile services/customer && \
	docker build -t admin:$(version) -f deployments/docker/app-go.Dockerfile services/admin && \
	docker build -t gateway:$(version) -f deployments/docker/app-go.Dockerfile services/api_gateway && \
	docker build -t order:$(version) -f deployments/docker/app-go.Dockerfile services/order && \
	docker build -t payment:$(version) -f deployments/docker/app-go.Dockerfile services/payment && \
	docker build -t product:$(version) -f deployments/docker/app-go.Dockerfile services/product && \
	minikube image load customer:$(version) && \
	minikube image load admin:$(version) && \
	minikube image load gateway:$(version) && \
	minikube image load order:$(version) && \
	minikube image load payment:$(version) && \
	minikube image load product:$(version) && \
	helm package deployments/k8s && \
	helm install ecommerce ./E-Commerce-0.1.0.tgz