package linsangmysql_test

//import (

//)

func ExampleIn() {

   hnd, err := New("mysqlhost:3306", "userid", "secret", "dbname")
   if err != nil {
      // Error handling stuff
   }

   defer hnd.Close()

   // Calling Create is only needed if you don't have a database yet, but it is not harmful
   err = hnd.Create()
   if err != nil {
      // Error handling stuff
   }

   // Calling Prepare is only needed after a call to New and you must have a initialized
   // database (if it is not the case, call Create before calling Prepare)
   err = hnd.Prepare()
   if err != nil {
      // Error handling stuff
   }

   // The queries below are done with data imported from
   // Cayley's (http://github.com/google/cayley/data/testdata.nq)
   // test data and follows the examples found at GREMLIN API
   // (https://github.com/google/cayley/blob/master/docs/GremlinAPI.md) documentation

   // Find the cool people, bob, greg and dani
   set, err = hnd.Walk(Vertex{"cool_person"},In{"status"})
   if err != nil {
      // Error handling stuff
   }


   // Find who follows bob, in this case, alice, charlie, and dani
   set, err = hnd.Walk(Vertex{"bob"},In{"follows"})
   if err != nil {
      // Error handling stuff
   }


   // Find who follows the people emily follows, namely, emily and bob
   set, err = hnd.Walk(Vertex{"emily"},Out{"follows"},In{"follows"})
   if err != nil {
      // Error handling stuff
   }

}

func ExampleOut() {

   hnd, err := New("mysqlhost:3306", "userid", "secret", "dbname")
   if err != nil {
      // Error handling stuff
   }

   defer hnd.Close()

   // Calling Create is only needed if you don't have a database yet, but it is not harmful
   err = hnd.Create()
   if err != nil {
      // Error handling stuff
   }

   // Calling Prepare is only needed after a call to New and you must have a initialized
   // database (if it is not the case, call Create before calling Prepare)
   err = hnd.Prepare()
   if err != nil {
      // Error handling stuff
   }

   // The queries below are done with data imported from
   // Cayley's (http://github.com/google/cayley/data/testdata.nq)
   // test data and follows the examples found at GREMLIN API
   // (https://github.com/google/cayley/blob/master/docs/GremlinAPI.md) documentation

   set, err = hnd.Walk(Vertex{"charlie"},Out{"follows"})
   if err != nil {
      // Error handling stuff
   }

   set, err = hnd.Walk(Vertex{"dani"},Out{"follows"},Out{"follows"})
   if err != nil {
      // Error handling stuff
   }

   set, err = hnd.Walk(Vertex{"dani"},Out{})
   if err != nil {
      // Error handling stuff
   }

   set, err = hnd.Walk(Vertex{"dani"},Out{[]string{"follows", "status"}})
   if err != nil {
      // Error handling stuff
   }

}

