async function test(): Promise<number> {
    return new Promise<number>(((resolve, reject) => {
        resolve(5);
    }));
}

function fib(n: number): number {
    return n < 2 ? 1 : fib(n - 2) + fib(n - 1)
}

async function fibAsync(n: number): Promise<number> {
    return new Promise(function (resolve, reject) {
        resolve(fib(n));
    })
}

async function singleThread(n:number): Promise<void> {
    console.log('计算斐波那契')
    let data = await fibAsync(n)
    console.log(data)
}

function singleThreadTest(){
    singleThread(30).then(()=>console.log("计算完毕"))
    console.log('计算中')
}
//await是一种替代异步回调的语法糖，代码看起来像同步，实际是异步执行
singleThreadTest()