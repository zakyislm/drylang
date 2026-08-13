
    use "dummy.y"
    
    e_IbRZdf = enc.b64("hello drylang")
    pt(e_IbRZdf)
    
    j_EztvTy = json(`{"test": 123}`)
    pt(j_EztvTy)
    
    // Test get() for type info
    val_KqDOqW = get("hello")
    pt(val_KqDOqW)
    

    fn calc_BqaEPg(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_LAqUxx = calc_BqaEPg(62, 90+1)
    pt(res_LAqUxx)
    

    sum_MzqfJK = 0
    lp 9 {
        sum_MzqfJK = sum_MzqfJK + 1
        if (sum_MzqfJK > 100) { done }
    }
    pt(sum_MzqfJK)
    

    m_LsjlES = math.sqrt(33)
    h_atKzAY = hash.md5("test_CgYQaF")
    log.inf("hash:", h_atKzAY)
    pt(m_LsjlES)
    

    m_wYxqIg = math.sqrt(48)
    h_ITOELQ = hash.md5("test_YqxOSP")
    log.inf("hash:", h_ITOELQ)
    pt(m_wYxqIg)
    

    try {
        throw_JRYxEl = unknown
        throw_JRYxEl()
    } err(e) {
        pt("caught error")
    }
    
