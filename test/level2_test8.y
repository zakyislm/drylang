
    cl Base_RKlDVk {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_ZeJAdi <- Base_RKlDVk {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_zSewTQ = Child_ZeJAdi(30, 34)
    pt(obj_zSewTQ.get_id())
    

    m_FgJkys = math.sqrt(30)
    h_sfTaXR = hash.md5("test_aQuxWu")
    log.inf("hash:", h_sfTaXR)
    pt(m_FgJkys)
    

    sum_rRSxKj = 0
    lp 9 {
        sum_rRSxKj = sum_rRSxKj + 1
        if (sum_rRSxKj > 100) { done }
    }
    pt(sum_rRSxKj)
    

    cl Base_YfKefj {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_wwLXPY <- Base_YfKefj {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ykYsHW = Child_wwLXPY(68, 24)
    pt(obj_ykYsHW.get_id())
    
