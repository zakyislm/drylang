
    use "dummy.y"
    
    e_BRKXOF = enc.b64("hello drylang")
    pt(e_BRKXOF)
    
    j_qctIGc = json(`{"test": 123}`)
    pt(j_qctIGc)
    
    // Test get() for type info
    val_gPyBPI = get("hello")
    pt(val_gPyBPI)
    

    try {
        throw_PFuqme = unknown
        throw_PFuqme()
    } err(e) {
        pt("caught error")
    }
    

    cns C_bRNyXa = 83
    v_KHDRvj = unknown
    v_KHDRvj = C_bRNyXa + 33
    if (v_KHDRvj > 0) {
        v_KHDRvj = v_KHDRvj * 2
    } el {
        v_KHDRvj = 0
    }
    pt(v_KHDRvj)
    

    m_shoiQe = math.sqrt(67)
    h_gHqmpz = hash.md5("test_gvybXs")
    log.inf("hash:", h_gHqmpz)
    pt(m_shoiQe)
    

    cns C_alYtLt = 90
    v_EZqyku = unknown
    v_EZqyku = C_alYtLt + 59
    if (v_EZqyku > 0) {
        v_EZqyku = v_EZqyku * 2
    } el {
        v_EZqyku = 0
    }
    pt(v_EZqyku)
    

    asn fn async_task_VRTpNk(x) {
        rev x * 2
    }
    uni async_task_VRTpNk(22)
    
    asn fn worker_FPZDcI(y) {
        pt("working on", y)
    }
    mul 2 worker_FPZDcI(56)
    
    awt
    

    cns C_DoHlhT = 16
    v_VGpXqH = unknown
    v_VGpXqH = C_DoHlhT + 91
    if (v_VGpXqH > 0) {
        v_VGpXqH = v_VGpXqH * 2
    } el {
        v_VGpXqH = 0
    }
    pt(v_VGpXqH)
    

    cns C_SXVDLm = 79
    v_wWeBoS = unknown
    v_wWeBoS = C_SXVDLm + 17
    if (v_wWeBoS > 0) {
        v_wWeBoS = v_wWeBoS * 2
    } el {
        v_wWeBoS = 0
    }
    pt(v_wWeBoS)
    

    arr_rMWgtD = [95, 68+1, 4+2]
    map_mGIEZC = {"a": arr_rMWgtD[0], "b": arr_rMWgtD[1]}
    pt(map_mGIEZC["a"])
    

    cns C_dHsyOc = 31
    v_GtrNdq = unknown
    v_GtrNdq = C_dHsyOc + 85
    if (v_GtrNdq > 0) {
        v_GtrNdq = v_GtrNdq * 2
    } el {
        v_GtrNdq = 0
    }
    pt(v_GtrNdq)
    

    m_TbwJdY = math.sqrt(59)
    h_ToiMnf = hash.md5("test_swumJP")
    log.inf("hash:", h_ToiMnf)
    pt(m_TbwJdY)
    

    fn calc_rUbtTf(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_nGcmpf = calc_rUbtTf(29, 100+1)
    pt(res_nGcmpf)
    

    fn calc_wouVtO(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_CMHfvR = calc_wouVtO(84, 26+1)
    pt(res_CMHfvR)
    

    arr_fFDxCN = [67, 79+1, 16+2]
    map_pkFBSB = {"a": arr_fFDxCN[0], "b": arr_fFDxCN[1]}
    pt(map_pkFBSB["a"])
    

    cl Base_PlVyuV {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_BPfGxU <- Base_PlVyuV {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_OYUqkN = Child_BPfGxU(10, 70)
    pt(obj_OYUqkN.get_id())
    

    arr_FlgJgn = [37, 87+1, 24+2]
    map_gwmrXm = {"a": arr_FlgJgn[0], "b": arr_FlgJgn[1]}
    pt(map_gwmrXm["a"])
    

    arr_imUVGk = [52, 55+1, 41+2]
    map_zCUIHR = {"a": arr_imUVGk[0], "b": arr_imUVGk[1]}
    pt(map_zCUIHR["a"])
    
