
    cl Base_xRzrAo {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_tuPdWW <- Base_xRzrAo {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_oHCqql = Child_tuPdWW(25, 52)
    pt(obj_oHCqql.get_id())
    

    cl Base_QOwlFB {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_YUsCvI <- Base_QOwlFB {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_cltDTq = Child_YUsCvI(67, 47)
    pt(obj_cltDTq.get_id())
    

    asn fn async_task_KKoshI(x) {
        rev x * 2
    }
    uni async_task_KKoshI(66)
    
    asn fn worker_kEciBq(y) {
        pt("working on", y)
    }
    mul 2 worker_kEciBq(6)
    
    awt
    

    fn calc_OWfxnY(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_fVnuSm = calc_OWfxnY(58, 23+1)
    pt(res_fVnuSm)
    

    asn fn async_task_NYlvzS(x) {
        rev x * 2
    }
    uni async_task_NYlvzS(76)
    
    asn fn worker_KnbdaL(y) {
        pt("working on", y)
    }
    mul 2 worker_KnbdaL(86)
    
    awt
    

    sum_OjxScI = 0
    lp 8 {
        sum_OjxScI = sum_OjxScI + 1
        if (sum_OjxScI > 100) { done }
    }
    pt(sum_OjxScI)
    

    val_PDAnGX = 88 % 3
    on (val_PDAnGX) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_PDAnGX)
    

    val_ureINM = 38 % 3
    on (val_ureINM) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_ureINM)
    
