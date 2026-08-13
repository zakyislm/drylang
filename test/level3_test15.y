
    use "dummy.y"
    
    e_GMcGwn = enc.b64("hello drylang")
    pt(e_GMcGwn)
    
    j_RkAAfZ = json(`{"test": 123}`)
    pt(j_RkAAfZ)
    
    // Test get() for type info
    val_xMXuqf = get("hello")
    pt(val_xMXuqf)
    

    fn calc_JUFVbH(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_yMWSFj = calc_JUFVbH(40, 33+1)
    pt(res_yMWSFj)
    

    fn calc_JUZyJx(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_yAnmZr = calc_JUZyJx(28, 62+1)
    pt(res_yAnmZr)
    

    cns C_wSQqcq = 75
    v_DNXWYc = unknown
    v_DNXWYc = C_wSQqcq + 1
    if (v_DNXWYc > 0) {
        v_DNXWYc = v_DNXWYc * 2
    } el {
        v_DNXWYc = 0
    }
    pt(v_DNXWYc)
    

    val_MsQMSC = 78 % 3
    on (val_MsQMSC) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_MsQMSC)
    
