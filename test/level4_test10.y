
    cns C_HwMkmC = 29
    v_NwMwkZ = unknown
    v_NwMwkZ = C_HwMkmC + 77
    if (v_NwMwkZ > 0) {
        v_NwMwkZ = v_NwMwkZ * 2
    } el {
        v_NwMwkZ = 0
    }
    pt(v_NwMwkZ)
    

    m_GbxFZK = math.sqrt(55)
    h_KdVqAR = hash.md5("test_hiHnuL")
    log.inf("hash:", h_KdVqAR)
    pt(m_GbxFZK)
    

    use "dummy.y"
    
    e_kuPKsY = enc.b64("hello drylang")
    pt(e_kuPKsY)
    
    j_ElnULU = json(`{"test": 123}`)
    pt(j_ElnULU)
    
    // Test get() for type info
    val_BZUUao = get("hello")
    pt(val_BZUUao)
    

    m_PHAYQh = math.sqrt(9)
    h_RVbtpz = hash.md5("test_QSSQdf")
    log.inf("hash:", h_RVbtpz)
    pt(m_PHAYQh)
    

    try {
        throw_nTCuOH = unknown
        throw_nTCuOH()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_yeFxNc = unknown
        throw_yeFxNc()
    } err(e) {
        pt("caught error")
    }
    
