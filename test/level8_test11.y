
    asn fn async_task_RETaHR(x) {
        rev x * 2
    }
    uni async_task_RETaHR(8)
    
    asn fn worker_EGBuLQ(y) {
        pt("working on", y)
    }
    mul 2 worker_EGBuLQ(20)
    
    awt
    

    sum_NtwQci = 0
    lp 7 {
        sum_NtwQci = sum_NtwQci + 1
        if (sum_NtwQci > 100) { done }
    }
    pt(sum_NtwQci)
    

    cns C_eNOvLm = 32
    v_KPXogz = unknown
    v_KPXogz = C_eNOvLm + 6
    if (v_KPXogz > 0) {
        v_KPXogz = v_KPXogz * 2
    } el {
        v_KPXogz = 0
    }
    pt(v_KPXogz)
    

    m_RqCPvo = math.sqrt(15)
    h_VEQdhn = hash.md5("test_owDePg")
    log.inf("hash:", h_VEQdhn)
    pt(m_RqCPvo)
    

    sum_feIMkY = 0
    lp 9 {
        sum_feIMkY = sum_feIMkY + 1
        if (sum_feIMkY > 100) { done }
    }
    pt(sum_feIMkY)
    

    cl Base_qYuwRX {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_QKqmYq <- Base_qYuwRX {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_tzZqpd = Child_QKqmYq(98, 35)
    pt(obj_tzZqpd.get_id())
    

    val_eIjacs = 82 % 3
    on (val_eIjacs) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_eIjacs)
    

    fn calc_FYpiFb(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_WUoFWL = calc_FYpiFb(42, 28+1)
    pt(res_WUoFWL)
    

    cl Base_jGubSk {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_tBGGEa <- Base_jGubSk {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_nzykzn = Child_tBGGEa(94, 60)
    pt(obj_nzykzn.get_id())
    

    fn calc_iBLamQ(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_cbxMXm = calc_iBLamQ(96, 21+1)
    pt(res_cbxMXm)
    

    m_OeqEsq = math.sqrt(21)
    h_LRDByD = hash.md5("test_uQHYmI")
    log.inf("hash:", h_LRDByD)
    pt(m_OeqEsq)
    

    cl Base_hUDTlF {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_HWgxDN <- Base_hUDTlF {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_JZJRWK = Child_HWgxDN(24, 61)
    pt(obj_JZJRWK.get_id())
    

    fn calc_TBrisc(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_ulDmcf = calc_TBrisc(74, 2+1)
    pt(res_ulDmcf)
    

    val_JswkTC = 7 % 3
    on (val_JswkTC) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_JswkTC)
    

    val_GAMlPm = 2 % 3
    on (val_GAMlPm) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_GAMlPm)
    
