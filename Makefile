# Rustライブラリのディレクトリ
RUST_LIB_DIR := calc_lib
# Goクライアントのディレクトリ
GO_CLIENT_DIR := calc_client
# ビルドされたRustライブラリのパス
RUST_LIB := $(RUST_LIB_DIR)/target/release/libcalc_lib.a
# Goクライアントの出力ファイル
GO_CLIENT := $(GO_CLIENT_DIR)/calc_client

# デフォルトターゲット
.PHONY: build
build: $(GO_CLIENT)

# Rustライブラリのビルドターゲット
$(RUST_LIB):
	@echo "Building Rust library..."
	cd $(RUST_LIB_DIR) && cargo build --release

# Goクライアントのビルドターゲット
$(GO_CLIENT): $(RUST_LIB)
	@echo "Copying Rust library to Go project..."
	mkdir -p $(GO_CLIENT_DIR)/lib
	cp $(RUST_LIB) $(GO_CLIENT_DIR)/lib/
	@echo "Building Go client..."
	cd $(GO_CLIENT_DIR) && go build -o calc_client -ldflags "-r $(GO_CLIENT_DIR)/lib"

# クリーンターゲット
.PHONY: clean
clean:
	@echo "Cleaning up..."
	cd $(RUST_LIB_DIR) && cargo clean
	rm -f $(GO_CLIENT)

# 実行ターゲット
.PHONY: run
run: $(GO_CLIENT)
	@echo "Running Go client..."
	cd $(GO_CLIENT_DIR) && ./calc_client
