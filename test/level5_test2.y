
    asn fn async_task_EhYxoW(x) {
        rev x * 2
    }
    uni async_task_EhYxoW(84)
    
    asn fn worker_xzmHhO(y) {
        pt("working on", y)
    }
    mul 2 worker_xzmHhO(76)
    
    awt
    

    cl Base_HqqpsO {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_blsirC <- Base_HqqpsO {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_cXwznv = Child_blsirC(68, 41)
    pt(obj_cXwznv.get_id())
    

    m_rskifH = math.sqrt(62)
    h_PAeNEk = hash.md5("test_HnmgUe")
    log.inf("hash:", h_PAeNEk)
    pt(m_rskifH)
    

    fn calc_rOrZQE(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_OScOrS = calc_rOrZQE(100, 93+1)
    pt(res_OScOrS)
    

    asn fn async_task_nnAoYd(x) {
        rev x * 2
    }
    uni async_task_nnAoYd(75)
    
    asn fn worker_ERNQyi(y) {
        pt("working on", y)
    }
    mul 2 worker_ERNQyi(45)
    
    awt
    

    asn fn async_task_EKEORH(x) {
        rev x * 2
    }
    uni async_task_EKEORH(40)
    
    asn fn worker_NiwBaX(y) {
        pt("working on", y)
    }
    mul 2 worker_NiwBaX(37)
    
    awt
    

    use "dummy.y"
    
    e_sIIgWY = enc.b64("hello drylang")
    pt(e_sIIgWY)
    
    j_NvtGbv = json(`{"test": 123}`)
    pt(j_NvtGbv)
    
    // Test get() for type info
    val_ohpPDw = get("hello")
    pt(val_ohpPDw)
    
