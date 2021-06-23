// 1. 运行Rust前必须先编译, 命令为 rustc 文件名
// 2. 如果是Windows平台, 还会额外生成一个pdb文件，其中包含调试信息
// 3. Rust编译出的程序, 可以在其他电脑运行，而无需安装Rust
// 4. rustc只适合简单的Rust程序, 复杂的代码则需要使用Cargo进行编译
// 5. 在安装Rust时将会默认安装Cargo
fn main() {
    // Rust的缩紧是四个空格，而不是tab
    println!("Hello World!");
}
