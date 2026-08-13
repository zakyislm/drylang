
    cns C_DkRUtN = 84
    v_IfnFQg = unknown
    v_IfnFQg = C_DkRUtN + 35
    if (v_IfnFQg > 0) {
        v_IfnFQg = v_IfnFQg * 2
    } el {
        v_IfnFQg = 0
    }
    pt(v_IfnFQg)
    

    use "dummy.y"
    
    e_gaCaRA = enc.b64("hello drylang")
    pt(e_gaCaRA)
    
    j_nsWPxs = json(`{"test": 123}`)
    pt(j_nsWPxs)
    
    // Test get() for type info
    val_XKccXR = get("hello")
    pt(val_XKccXR)
    

    m_nFjHIO = math.sqrt(27)
    h_erVZRK = hash.md5("test_KgItbB")
    log.inf("hash:", h_erVZRK)
    pt(m_nFjHIO)
    

    cns C_UJBCdr = 100
    v_zOTUzf = unknown
    v_zOTUzf = C_UJBCdr + 50
    if (v_zOTUzf > 0) {
        v_zOTUzf = v_zOTUzf * 2
    } el {
        v_zOTUzf = 0
    }
    pt(v_zOTUzf)
    

    fn calc_xFwIxt(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_feSTqh = calc_xFwIxt(87, 52+1)
    pt(res_feSTqh)
    

    cns C_aljkWW = 70
    v_EbAzNI = unknown
    v_EbAzNI = C_aljkWW + 18
    if (v_EbAzNI > 0) {
        v_EbAzNI = v_EbAzNI * 2
    } el {
        v_EbAzNI = 0
    }
    pt(v_EbAzNI)
    

    sum_WXnaJS = 0
    lp 7 {
        sum_WXnaJS = sum_WXnaJS + 1
        if (sum_WXnaJS > 100) { done }
    }
    pt(sum_WXnaJS)
    
