package cache

// TODO tcp模式
//import (
//	"bufio"
//	"io"
//	"log"
//	"net"
//)
//
//type TcpHandler struct {
//	cache *Cache
//}
//
//func (th *TcpHandler) Listen(port string) {
//	l, err := net.Listen("tcp", "127.0.0.1:"+port)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	for {
//		conn, err := l.Accept()
//		if err != nil {
//			log.Fatalln(err)
//		}
//		go th.process(conn)
//	}
//
//}
//
//func New(c *Cache) *TcpHandler {
//	return &TcpHandler{cache: c}
//}
//
//type InterfaceCache interface {
//	SetString(string, string, int64) bool
//	GetString(string) *string
//	RmString(string) bool
//}
//
//type result struct {
//	v   *string
//	err error
//}
//
//func (th *TcpHandler) process(conn net.Conn) {
//	defer conn.Close()
//	r := bufio.NewReader(conn)
//	resultCh := make(chan chan *result, 1000)
//	go reply(conn, resultCh)
//
//	for {
//		op, err := r.ReadByte()
//		if err != nil {
//			if err != io.EOF {
//				log.Println("close the conn due to error:", err)
//			}
//			return
//		}
//
//		if op == 'S' {
//			th.set(resultCh, r)
//		} else if op == 'G' {
//			th.get(resultCh, r)
//		} else if op == 'D' {
//			th.del(resultCh, r)
//		} else {
//			log.Println("close the conn due to invalid op:", op)
//			return
//		}
//	}
//}
//
//func (th *TcpHandler) get(ch chan chan *result, r *bufio.Reader) {
//	c := make(chan *result)
//	ch <- c
//
//	k, err := th.readKey(r)
//	if err != nil {
//		c <- &result{nil, err}
//		return
//	}
//
//	go func() {
//		v := th.cache.GetString(k)
//		c <- &result{v, err}
//	}()
//
//}
//
//func (th *TcpHandler) set(ch chan chan *result, r *bufio.Reader) {
//	c := make(chan *result)
//	ch <- c
//
//	k, v, err := th.readKeyAndValue(r)
//	if err != nil {
//		c <- &result{nil, err}
//		return
//	}
//
//	go func() {
//		c <- &result{nil, th.cache.SetString(k, v)}
//	}()
//
//}
//
//func (th *TcpHandler) del(ch chan chan *result, r *bufio.Reader) {
//	c := make(chan *result)
//	ch <- c
//
//	k, err := th.readKey(r)
//	if err != nil {
//		c <- &result{nil, err}
//		return
//	}
//
//	go func() {
//		c <- &result{nil, th.Del(k)}
//	}()
//
//}
//
//func (th *TcpHandler) readKey(r *bufio.Reader) (string, error) {
//	klen, e := readLen(r)
//	if e != nil {
//		return "", e
//	}
//	k := make([]byte, klen)
//	_, e = io.ReadFull(r, k)
//	if e != nil {
//		return "", e
//	}
//	return string(k), nil
//}
//
//func (th *TcpHandler) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
//	klen, e := readLen(r)
//	if e != nil {
//		return "", nil, e
//	}
//
//	vlen, e := readLen(r)
//	if e != nil {
//		return "", nil, e
//	}
//
//	k := make([]byte, klen)
//	_, e = io.ReadFull(r, k)
//	if e != nil {
//		return "", nil, e
//	}
//
//	v := make([]byte, vlen)
//	_, e = io.ReadFull(r, v)
//	if e != nil {
//		return "", nil, e
//	}
//
//	return string(k), v, nil
//
//}
//
//func reply(conn net.Conn, resultCh chan chan *result) {
//	defer conn.Close()
//
//	for {
//		c, open := <-resultCh
//		if !open {
//			return
//		}
//
//		r := <-c
//		err := sendResponse(r.v, r.err, conn)
//		if err != nil {
//			log.Println("close the conn due to error:", err)
//			return
//		}
//	}
//}
