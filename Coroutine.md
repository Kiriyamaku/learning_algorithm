
# 协程

## 是什么
同进程内的一种用户态调度机制

协程看上去也是子程序，但执行过程中，在子程序内部可中断，然后转而执行别的子程序，在适当的时候再返回来接着执行。

注意，在一个子程序中中断，去执行其他子程序，不是函数调用，有点类似CPU的**中断**。
## 优势
- 不需要锁，因为同时间只有一段代码运行
- 开销小（相对于多线程来说），单进程内切换
- 各个协程之间的切换，往往是用户通过代码来显式指定的（跟各种 callback 类似），不需要内核参与，可以很方便的实现异步。
## 缺陷
- 缺点是如果其中一个协程中有密集计算，其他的协程就不运行了。操作系统进程的缺点是开销大，优点是无论代码怎么写，所有进程都可以并发运行。

## 应用
- 许多语言支持协程

### 解决问题
- 用一组少量的线程来实现多个任务，一旦某个任务阻塞，则可能用同一线程继续运行其他任务，避免大量上下文的切换
## 例子（python）
> from [廖雪峰](https://www.liaoxuefeng.com/wiki/897692888725344/923057403198272)
```py
import time

def consumer():
    r = ''
    while True:
        n = yield r
        if not n:
            # 
            return 
        print('[CONSUMER] Consuming %s...' % n)
        time.sleep(1)
        r = '200 OK'

def produce(c):

    # 首先调用c.next()启动生成器；
    c.next() 
    n = 0
    while n < 5:
        n = n + 1
        print('[PRODUCER] Producing %s...' % n)
        # 然后，一旦生产了东西，通过c.send(n)切换到consumer执行；
        r = c.send(n)
        print('[PRODUCER] Consumer return: %s' % r)
    c.close()

if __name__=='__main__':
    c = consumer()
    produce(c)
```
result:
```sh
[PRODUCER] Producing 1...
[CONSUMER] Consuming 1...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 2...
[CONSUMER] Consuming 2...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 3...
[CONSUMER] Consuming 3...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 4...
[CONSUMER] Consuming 4...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 5...
[CONSUMER] Consuming 5...
[PRODUCER] Consumer return: 200 OK
```