package bench

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Perform some work
func work() {
	for i := 0; i < 10; i++ {
		_ = 1 + 1
	}
}

func BenchmarkLocks(b *testing.B) {
	b.Run("Mutex", func(b *testing.B) {
		m := sync.Mutex{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Lock()
				work()
				m.Unlock()
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					m.Lock()
					work()
					m.Unlock()
				}
			})
		})
	})
	b.Run("RWMutex", func(b *testing.B) {
		m := sync.RWMutex{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Lock()
				work()
				m.Unlock()
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					m.Lock()
					work()
					m.Unlock()
				}
			})
		})
	})
	b.Run("RWMutexRead", func(b *testing.B) {
		m := sync.RWMutex{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.RLock()
				work()
				m.RUnlock()
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					m.RLock()
					work()
					m.RUnlock()
				}
			})
		})
	})
	b.Run("Bool", func(b *testing.B) {
		lk := atomic.Bool{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !lk.CompareAndSwap(false, true) {
				} // Lock
				work()
				lk.Store(false) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !lk.CompareAndSwap(false, true) {
					} // Lock
					work()
					lk.Store(false) // Unlock
				}
			})
		})
	})
	b.Run("int32", func(b *testing.B) {
		var lk int32 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapInt32(&lk, 0, 1) {
				} // Lock
				work()
				atomic.StoreInt32(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapInt32(&lk, 0, 1) {
					} // Lock
					work()
					atomic.StoreInt32(&lk, 0) // Unlock
				}
			})
		})
	})
	b.Run("int64", func(b *testing.B) {
		var lk int64 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapInt64(&lk, 0, 1) {
				} // Lock
				work()
				atomic.StoreInt64(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapInt64(&lk, 0, 1) {
					} // Lock
					work()
					atomic.StoreInt64(&lk, 0) // Unlock
				}
			})
		})
	})
	b.Run("Uint32", func(b *testing.B) {
		var lk uint32 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapUint32(&lk, 0, 1) {
				} // Lock
				work()
				atomic.StoreUint32(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapUint32(&lk, 0, 1) {
					} // Lock
					work()
					atomic.StoreUint32(&lk, 0) // Unlock
				}
			})
		})
	})
	b.Run("Uint64", func(b *testing.B) {
		var lk uint64 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapUint64(&lk, 0, 1) {
				} // Lock
				work()
				atomic.StoreUint64(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapUint64(&lk, 0, 1) {
					} // Lock
					work()
					atomic.StoreUint64(&lk, 0) // Unlock
				}
			})
		})
	})
}

func BenchmarkLocksNoWork(b *testing.B) {
	b.Run("Mutex", func(b *testing.B) {
		m := sync.Mutex{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Lock()
				m.Unlock()
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					m.Lock()
					m.Unlock()
				}
			})
		})
	})
	b.Run("RWMutex", func(b *testing.B) {
		m := sync.RWMutex{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Lock()
				m.Unlock()
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					m.Lock()
					m.Unlock()
				}
			})
		})
	})
	b.Run("RWMutexRead", func(b *testing.B) {
		m := sync.RWMutex{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.RLock()
				m.RUnlock()
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					m.RLock()
					m.RUnlock()
				}
			})
		})
	})
	b.Run("Bool", func(b *testing.B) {
		lk := atomic.Bool{}
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !lk.CompareAndSwap(false, true) {
				} // Lock
				lk.Store(false) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !lk.CompareAndSwap(false, true) {
					} // Lock
					lk.Store(false) // Unlock
				}
			})
		})
	})
	b.Run("int32", func(b *testing.B) {
		var lk int32 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapInt32(&lk, 0, 1) {
				} // Lock
				atomic.StoreInt32(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapInt32(&lk, 0, 1) {
					} // Lock
					atomic.StoreInt32(&lk, 0) // Unlock
				}
			})
		})
	})
	b.Run("int64", func(b *testing.B) {
		var lk int64 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapInt64(&lk, 0, 1) {
				} // Lock
				atomic.StoreInt64(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapInt64(&lk, 0, 1) {
					} // Lock
					atomic.StoreInt64(&lk, 0) // Unlock
				}
			})
		})
	})
	b.Run("Uint32", func(b *testing.B) {
		var lk uint32 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapUint32(&lk, 0, 1) {
				} // Lock
				atomic.StoreUint32(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapUint32(&lk, 0, 1) {
					} // Lock
					atomic.StoreUint32(&lk, 0) // Unlock
				}
			})
		})
	})
	b.Run("Uint64", func(b *testing.B) {
		var lk uint64 = 0
		b.Run("Sequential", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for !atomic.CompareAndSwapUint64(&lk, 0, 1) {
				} // Lock
				atomic.StoreUint64(&lk, 0) // Unlock
			}
		})
		b.Run("Parallel", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for !atomic.CompareAndSwapUint64(&lk, 0, 1) {
					} // Lock
					atomic.StoreUint64(&lk, 0) // Unlock
				}
			})
		})
	})
}
