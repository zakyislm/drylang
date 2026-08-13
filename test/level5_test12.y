
    cns C_PQWiGt = 56
    v_hdFmaw = unknown
    v_hdFmaw = C_PQWiGt + 72
    if (v_hdFmaw > 0) {
        v_hdFmaw = v_hdFmaw * 2
    } el {
        v_hdFmaw = 0
    }
    pt(v_hdFmaw)
    

    cns C_fnENSB = 3
    v_EWGpGY = unknown
    v_EWGpGY = C_fnENSB + 80
    if (v_EWGpGY > 0) {
        v_EWGpGY = v_EWGpGY * 2
    } el {
        v_EWGpGY = 0
    }
    pt(v_EWGpGY)
    

    sum_sGksdE = 0
    lp 12 {
        sum_sGksdE = sum_sGksdE + 1
        if (sum_sGksdE > 100) { done }
    }
    pt(sum_sGksdE)
    

    sum_pdmFLR = 0
    lp 14 {
        sum_pdmFLR = sum_pdmFLR + 1
        if (sum_pdmFLR > 100) { done }
    }
    pt(sum_pdmFLR)
    

    arr_rdVHag = [98, 94+1, 99+2]
    map_hxmlWp = {"a": arr_rdVHag[0], "b": arr_rdVHag[1]}
    pt(map_hxmlWp["a"])
    

    m_REjTUP = math.sqrt(83)
    h_YverqZ = hash.md5("test_jvyiMt")
    log.inf("hash:", h_YverqZ)
    pt(m_REjTUP)
    

    val_RygFsh = 17 % 3
    on (val_RygFsh) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_RygFsh)
    
