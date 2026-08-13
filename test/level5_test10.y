
    use "dummy.y"
    
    e_txaoNN = enc.b64("hello drylang")
    pt(e_txaoNN)
    
    j_imrkkH = json(`{"test": 123}`)
    pt(j_imrkkH)
    
    // Test get() for type info
    val_rJULpr = get("hello")
    pt(val_rJULpr)
    

    try {
        throw_bUEwoT = unknown
        throw_bUEwoT()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_uqUCKz(x) {
        rev x * 2
    }
    uni async_task_uqUCKz(65)
    
    asn fn worker_OatveY(y) {
        pt("working on", y)
    }
    mul 2 worker_OatveY(52)
    
    awt
    

    m_WiuPff = math.sqrt(46)
    h_NOzSfd = hash.md5("test_vPiRFe")
    log.inf("hash:", h_NOzSfd)
    pt(m_WiuPff)
    

    sum_KaDqzT = 0
    lp 10 {
        sum_KaDqzT = sum_KaDqzT + 1
        if (sum_KaDqzT > 100) { done }
    }
    pt(sum_KaDqzT)
    

    arr_VIGMOy = [100, 14+1, 75+2]
    map_gCAnXc = {"a": arr_VIGMOy[0], "b": arr_VIGMOy[1]}
    pt(map_gCAnXc["a"])
    

    try {
        throw_UDxTpZ = unknown
        throw_UDxTpZ()
    } err(e) {
        pt("caught error")
    }
    
