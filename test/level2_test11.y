
    fn calc_UGnZSi(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_mVIODi = calc_UGnZSi(16, 69+1)
    pt(res_mVIODi)
    

    val_WlzOvx = 79 % 3
    on (val_WlzOvx) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_WlzOvx)
    

    m_qbRLwx = math.sqrt(9)
    h_mrKJPQ = hash.md5("test_asKvzQ")
    log.inf("hash:", h_mrKJPQ)
    pt(m_qbRLwx)
    

    cl Base_ujxwIQ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_srhgsp <- Base_ujxwIQ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_UDYgpd = Child_srhgsp(20, 28)
    pt(obj_UDYgpd.get_id())
    
