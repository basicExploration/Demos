package main

r := gin.Default()
store := cookie.NewStore([]byte("testCookie")) // 这个就是cookie的名字而已。
r.Use(sessions.Sessions("mysession", store))   // 这个就是 将这个store传递给了sessions中然后再选个mysession然后就ok了。
r.GET("/", func(c *gin.Context) {
	session := sessions.Default(c) //  sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count) /// set count 就是在这个sessions的session的value。
	session.Save()
	c.JSON(200, gin.H{"count": count})
})
r.Run(":8000")
