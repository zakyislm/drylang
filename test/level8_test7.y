
    asn fn async_task_aUkDPB(x) {
        rev x * 2
    }
    uni async_task_aUkDPB(11)
    
    asn fn worker_MYYPWq(y) {
        pt("working on", y)
    }
    mul 2 worker_MYYPWq(93)
    
    awt
    

    m_mFNGDz = math.sqrt(3)
    h_YnszrM = hash.md5("test_ovtIrY")
    log.inf("hash:", h_YnszrM)
    pt(m_mFNGDz)
    

    arr_sVkBVG = [16, 60+1, 23+2]
    map_zooaTH = {"a": arr_sVkBVG[0], "b": arr_sVkBVG[1]}
    pt(map_zooaTH["a"])
    

    sum_SvWlLU = 0
    lp 13 {
        sum_SvWlLU = sum_SvWlLU + 1
        if (sum_SvWlLU > 100) { done }
    }
    pt(sum_SvWlLU)
    

    use "dummy.y"
    
    e_GRvnQV = enc.b64("hello drylang")
    pt(e_GRvnQV)
    
    j_qoIDJF = json(`{"test": 123}`)
    pt(j_qoIDJF)
    
    // Test get() for type info
    val_kxrojq = get("hello")
    pt(val_kxrojq)
    

    sum_yZllhY = 0
    lp 6 {
        sum_yZllhY = sum_yZllhY + 1
        if (sum_yZllhY > 100) { done }
    }
    pt(sum_yZllhY)
    

    asn fn async_task_woNnfa(x) {
        rev x * 2
    }
    uni async_task_woNnfa(65)
    
    asn fn worker_mUmvpU(y) {
        pt("working on", y)
    }
    mul 2 worker_mUmvpU(63)
    
    awt
    

    use "dummy.y"
    
    e_nOpXPp = enc.b64("hello drylang")
    pt(e_nOpXPp)
    
    j_tyhkgw = json(`{"test": 123}`)
    pt(j_tyhkgw)
    
    // Test get() for type info
    val_tyMwTw = get("hello")
    pt(val_tyMwTw)
    

    m_XQTMfP = math.sqrt(24)
    h_NNiuqy = hash.md5("test_ARnbwu")
    log.inf("hash:", h_NNiuqy)
    pt(m_XQTMfP)
    

    fn calc_HcIjGM(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_KYrdpp = calc_HcIjGM(24, 56+1)
    pt(res_KYrdpp)
    

    arr_lazpit = [74, 31+1, 12+2]
    map_xAoagP = {"a": arr_lazpit[0], "b": arr_lazpit[1]}
    pt(map_xAoagP["a"])
    

    val_vPlhof = 38 % 3
    on (val_vPlhof) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_vPlhof)
    

    use "dummy.y"
    
    e_EVQWtJ = enc.b64("hello drylang")
    pt(e_EVQWtJ)
    
    j_wDTZwF = json(`{"test": 123}`)
    pt(j_wDTZwF)
    
    // Test get() for type info
    val_rJcEsg = get("hello")
    pt(val_rJcEsg)
    

    fn calc_IOCwKg(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_Xbpxkk = calc_IOCwKg(32, 80+1)
    pt(res_Xbpxkk)
    

    cl Base_ZdcVtP {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CBwbaN <- Base_ZdcVtP {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_QXfUzI = Child_CBwbaN(17, 75)
    pt(obj_QXfUzI.get_id())
    
