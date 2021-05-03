# js

## symbol

Symbol用于防止属性名冲突而产生的，比如向第三方对象中添加属性时。

使用`description`可以获取传入的描述参数

Symbol 声明和访问使用 `[]`（变量）形式操作

```js
let symbol = Symbol("后盾人");
let obj = {
  symbol: "hdcms.com"
};
console.log(obj); //wrong

let symbol = Symbol("后盾人");
let obj = {
  [symbol]: "houdunren.com"
};//right
```



如果对象属性不想被遍历，可以使用`Symbol`保护

```
const site = Symbol("网站名称");
class User {
  constructor(name) {
    this[site] = "后盾人";
    this.name = name;
  }
  getName() {
    return `${this[site]}-${this.name}`;
  }
}
const hd = new User("向军大叔");
console.log(hd.getName());
for (const key in hd) {
  console.log(key);
}
```

## set

```js
let obj = { 1: "hdcms", "1": "houdunren" };
console.table(obj); //{1:"houdunren"}
```

```js
let set = new Set();
set.add(1);
set.add("1");
console.log(set); //Set(2) {1, "1"}
```

方法：

delete

clear

add

has

Array.from(set) : 将set转换为array

```js
console.log([...set]); //["hdcms", "houdunren"]
```



使用场景：数组去重

set只有值，因此value和key相同

```js
let arr = [7, 6, 2, 8, 2, 6];
let set = new Set(arr);
//使用forEach遍历
set.forEach((item,key) => console.log(item,key));
```



遍历：forEach

## map

```js
let hd = new Map([["houdunren", "后盾人"], ["hdcms", "开源系统"]]);

console.log(...hd); //(2) ["houdunren", "后盾人"] (2) ["hdcms", "开源系统"]
console.log(...hd.entries()); //(2) ["houdunren", "后盾人"] (2) ["hdcms", "开源系统"]
console.log(...hd.values()); //后盾人 开源系统
console.log(...hd.keys()); //houdunren hdcms
```

## function

arguments:得到所有参数的集合

```
function sum() {
  return [...arguments].reduce((total, num) => {
    return (total += num);
  }, 0);
}
console.log(sum(2, 3, 4, 2, 6)); //17
```

也可以

```js
function sum(...args) {
 return args.reduce((a, b) => a + b);
}
console.log(sum(2, 3, 4, 2, 6)); //17
```

### call/apply

call与apply 用于显示的设置函数的上下文，两个方法作用一样都是将对象绑定到this，只是在传递参数上有所不同。

- apply 用数组传参
- call 需要分别传参
- 与 bind 不同 call/apply 会立即执行函数

```js
function show(title) {
    alert(`${title+this.name}`);
}
let lisi = {
    name: '李四'
};
let wangwu = {
    name: '王五'
};
show.call(lisi, '后盾人');
show.apply(wangwu, ['HDCMS']);
```

### bind

bind()是将函数绑定到某个对象，比如 a.bind(hd) 可以理解为将a函数绑定到hd对象上即 hd.a()。

- 与 call/apply 不同bind不会立即执行
- bind 是复制函数形为会返回新函数

```js
function hd(a, b) {
  return this.f + a + b;
}

//使用bind会生成新函数
let newFunc = hd.bind({ f: 1 }, 3);

//1+3+2 参数2赋值给b即 a=3,b=2
console.log(newFunc(2));
```



## closure

当一个变量既不是该函数内部的局部变量,也不是该函数的参数,相对于作用域来说,就是一个自由变量(引用了外部变量),这样就会形成一个闭包.



## promise

> promise 是一个拥有 `then` 方法的对象或函数

`JavaScript` 中存在很多异步操作,`Promise` 将异步操作队列化，按照期望的顺序执行，返回符合预期的结果。可以通过链式调用多个 `Promise` 达到我们的目的。

Promise 可以理解为承诺，就像我们去KFC点餐服务员给我们一引取餐票，这就是承诺。如果餐做好了叫我们这就是成功，如果没有办法给我们做出食物这就是拒绝。

- 一个 `promise` 必须有一个 `then` 方法用于处理状态改变

Promise包含`pending`、`fulfilled`、`rejected`三种状态

- `pending` 指初始等待状态，初始化 `promise` 时的状态
- `resolve` 指已经解决，将 `promise` 状态设置为`fulfilled`
- `reject` 指拒绝处理，将 `promise` 状态设置为`rejected`
- `promise` 是生产者，通过 `resolve` 与 `reject` 函数告之结果
- `promise` 非常适合需要一定执行时间的异步任务
- 状态一旦改变将不可更改

`promise` 创建时即立即执行即同步任务，`then` 会放在异步微任务中执行，需要等同步任务执行后才执行。

```js
const pro = new Promise((resolve, reject) =>{
    // resolve("fulfilled!");
    reject("reject")
})
console.log(pro);
```

- `promise` 的 then、catch、finally的方法都是异步任务
- 程序需要将主任务执行完成才会执行异步队列任务

### then

一个promise 需要提供一个then方法访问promise 结果，`then` 用于定义当 `promise` 状态发生改变时的处理，即`promise`处理异步操作，`then` 用于结果。

`promise` 就像 `kfc` 中的厨房，`then` 就是我们用户，如果餐做好了即 `fulfilled` ，做不了拒绝即`rejected` 状态。那么 then 就要对不同状态处理。

- then 方法必须返回 promise，用户返回或系统自动返回
- 第一个函数在`resolved` 状态时执行，即执行`resolve`时执行`then`第一个函数处理成功状态
- 第二个函数在`rejected`状态时执行，即执行`reject` 时执行第二个函数处理失败状态，该函数是可选的
- 两个函数都接收 `promise` 传出的值做为参数
- 也可以使用`catch` 来处理失败的状态
- 如果 `then` 返回 `promise` ，下一个`then` 会在当前`promise` 状态改变后执行

```js
promise.then(onFulfilled, onRejected)
```

```js
const promise = new Promise((resolve, reject) => {
  resolve("success");
}).then(
  value => {
    console.log(`解决：${value}`);//onFulfilled
  },
  reason => {
    console.log(`拒绝:${reason}`);//onRejected
  }
);
```

两个都是函数

每次的 `then` 都是一个全新的 `promise`，默认 then 返回的 promise 状态是 fulfilled

每次的 `then` 都是一个全新的 `promise`，不要认为上一个 promise状态会影响以后then返回的状态

###  async/await

使用 `async/await` 是promise 的语法糖，可以让编写 promise 更清晰易懂，也是推荐编写promise 的方式。

- `async/await` 本质还是promise，只是更简洁的语法糖书写
- `async/await`  使用更清晰的promise来替换 promise.then/catch 的方式



下面在 `hd` 函数前加上async，函数将返回promise，我们就可以像使用标准Promise一样使用了。

### 

使用 `await` 关键词后会等待promise 完

- `await` 后面一般是promise，如果不是直接返回
- `await` 必须放在 async 定义的函数中使用
- `await` 用于替代 `then` 使编码更优雅