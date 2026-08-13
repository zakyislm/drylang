
    asn fn async_task_bkUBSe(x) {
        rev x * 2
    }
    uni async_task_bkUBSe(54)
    
    asn fn worker_kwNTgK(y) {
        pt("working on", y)
    }
    mul 2 worker_kwNTgK(84)
    
    awt
    

    m_oJTbZu = math.sqrt(19)
    h_QnBdsl = hash.md5("test_mpOIKC")
    log.inf("hash:", h_QnBdsl)
    pt(m_oJTbZu)
    

    m_viivZW = math.sqrt(63)
    h_FcFKZm = hash.md5("test_rqkdja")
    log.inf("hash:", h_FcFKZm)
    pt(m_viivZW)
    

    arr_MOGzdv = [76, 27+1, 94+2]
    map_OSLUmw = {"a": arr_MOGzdv[0], "b": arr_MOGzdv[1]}
    pt(map_OSLUmw["a"])
    

    m_cOwxEv = math.sqrt(81)
    h_iUqcCd = hash.md5("test_pZoNdk")
    log.inf("hash:", h_iUqcCd)
    pt(m_cOwxEv)
    

    fn calc_xuryqG(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_edsKEu = calc_xuryqG(86, 42+1)
    pt(res_edsKEu)
    

    arr_hwjrxH = [53, 51+1, 46+2]
    map_KCbMPn = {"a": arr_hwjrxH[0], "b": arr_hwjrxH[1]}
    pt(map_KCbMPn["a"])
    

    use "dummy.y"
    
    e_VHNGqW = enc.b64("hello drylang")
    pt(e_VHNGqW)
    
    j_XRKRme = json(`{"test": 123}`)
    pt(j_XRKRme)
    
    // Test get() for type info
    val_MmuwxM = get("hello")
    pt(val_MmuwxM)
    

    cns C_jbiYai = 1
    v_tooQpP = unknown
    v_tooQpP = C_jbiYai + 88
    if (v_tooQpP > 0) {
        v_tooQpP = v_tooQpP * 2
    } el {
        v_tooQpP = 0
    }
    pt(v_tooQpP)
    

    arr_LeqJOq = [62, 47+1, 22+2]
    map_jYlbVR = {"a": arr_LeqJOq[0], "b": arr_LeqJOq[1]}
    pt(map_jYlbVR["a"])
    

    cl Base_zSktlu {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CUMEMT <- Base_zSktlu {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_PWpitT = Child_CUMEMT(14, 90)
    pt(obj_PWpitT.get_id())
    

    fn calc_bLmECD(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_jyikef = calc_bLmECD(10, 33+1)
    pt(res_jyikef)
    

    arr_zFZDlf = [87, 16+1, 91+2]
    map_vQXmwn = {"a": arr_zFZDlf[0], "b": arr_zFZDlf[1]}
    pt(map_vQXmwn["a"])
    

    fn calc_AlvAzq(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_uefvfe = calc_AlvAzq(27, 52+1)
    pt(res_uefvfe)
    
