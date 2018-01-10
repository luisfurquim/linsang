package linsangmysql

/*
Executes initialization procedures
*/
func (hnd  *Handle) Open() error {
   var err error

   err = hnd.prepare()
   if err != nil {
      err = hnd.Create()
      if err != nil {
         Goose.Init.Logf(1,"Error creating DB: %s\n",err)
         return err
      }
   }

   err = hnd.prepare()
   if err != nil {
      Goose.Init.Logf(1,"Error preparing DB: %s\n",err)
      return err
   }

   return nil
}
