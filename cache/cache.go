package cache2go

import (
  "sync"
)

var (
  cache = make(map[string]*CacheTable)
  // RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
  mutex sync.RWMutex
)

func Cache(table string) *CacheTable  {
  // Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入。
  mutex.RLock()
  t, ok := cache[table]
  // Runlock方法解除rw的读取锁状态，如果m未加读取锁会导致运行时错误。
  mutex.RUnlock()

  if !ok {
    mutex.Lock()
    t, ok = cache[table]
    if !ok {
      t = &CacheTable{
        name: table,
        items: make(map[interface{}]*CacheItem)
      }
      cache[table] = t
    }
    // Unlock方法解除rw的写入锁状态，如果m未加写入锁会导致运行时错误。
    mutex.Unlock()
  }
  return t
}
