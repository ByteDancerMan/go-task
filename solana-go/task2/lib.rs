use anchor_lang::prelude::*; // 导入Anchor框架的所有必要组件

// This is your program's public key and it will update
// automatically when you build the project.
declare_id!("11111111111111111111111111111111"); // 声明程序的公共密钥（Program ID），在构建时会自动更新为实际地址

#[program] // 属性宏，标记这是程序的主逻辑模块
mod hello_anchor { // 定义模块名称
    use super::*;
    //传入的参数，类型为64位无符号整数
    pub fn initialize(ctx: Context<Initialize>, data: u64) -> Result<()> { // 定义公共函数initialize，作为程序的一个可调用指令
        ctx.accounts.new_account.data = data; // 将传入的数据存储到账户中
        msg!("Changed data to: {}!", data); // 记录日志消息，会在交易日志中显示
        Ok(()) // 返回成功结果
    }
}

#[derive(Accounts)] // 派生宏，自动生成账户处理逻辑
pub struct Initialize<'info> {
    // We must specify the space in order to initialize an account.
    // First 8 bytes are default account discriminator,
    // next 8 bytes come from NewAccount.data being type u64.
    // (u64 = 64 bits unsigned integer = 8 bytes)
    #[account(
      init, // 表示要初始化一个新账户
      payer = signer, // 账户的拥有者
      space = 8 + 8 // 账户所需的存储空间（字节数）
    )]
    pub new_account: Account<'info, NewAccount>, // 账户结构名称
    #[account(mut)] // 账户可变
    pub signer: Signer<'info>, // 账户拥有者
    pub system_program: Program<'info, System>, // 系统程序
}

#[account] // 派生宏，自动生成账户结构处理逻辑
pub struct NewAccount { // 定义账户结构名称
    data: u64 // 定义账户结构，包含一个64位无符号整数字段
}