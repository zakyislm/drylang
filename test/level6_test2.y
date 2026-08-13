
    m_NqbnKd = math.sqrt(70)
    h_EatZJu = hash.md5("test_SDMGPw")
    log.inf("hash:", h_EatZJu)
    pt(m_NqbnKd)
    

    use "dummy.y"
    
    e_ZBdcrH = enc.b64("hello drylang")
    pt(e_ZBdcrH)
    
    j_lqmSek = json(`{"test": 123}`)
    pt(j_lqmSek)
    
    // Test get() for type info
    val_zasUQI = get("hello")
    pt(val_zasUQI)
    

    arr_oDyOfY = [2, 66+1, 97+2]
    map_FQjooj = {"a": arr_oDyOfY[0], "b": arr_oDyOfY[1]}
    pt(map_FQjooj["a"])
    

    val_gUuRbl = 14 % 3
    on (val_gUuRbl) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_gUuRbl)
    

    use "dummy.y"
    
    e_RnqjSu = enc.b64("hello drylang")
    pt(e_RnqjSu)
    
    j_nsIBWn = json(`{"test": 123}`)
    pt(j_nsIBWn)
    
    // Test get() for type info
    val_iQXizs = get("hello")
    pt(val_iQXizs)
    

    m_mhqzQh = math.sqrt(50)
    h_yvWbBN = hash.md5("test_MoUroS")
    log.inf("hash:", h_yvWbBN)
    pt(m_mhqzQh)
    

    sum_NaoWMr = 0
    lp 13 {
        sum_NaoWMr = sum_NaoWMr + 1
        if (sum_NaoWMr > 100) { done }
    }
    pt(sum_NaoWMr)
    

    val_WOSYwd = 62 % 3
    on (val_WOSYwd) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_WOSYwd)
    
