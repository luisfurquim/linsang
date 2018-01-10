package linsangmysql

// Closes the MySQL storage
func (hnd *Handle) Close() error {
   return hnd.Conn.Close()
}
