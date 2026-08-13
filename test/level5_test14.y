
    val_eETlwf = 88 % 3
    on (val_eETlwf) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_eETlwf)
    

    cl Base_RPTtea {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_PIAtuX <- Base_RPTtea {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_hrwiZY = Child_PIAtuX(79, 51)
    pt(obj_hrwiZY.get_id())
    

    m_ZQoBcE = math.sqrt(24)
    h_eLOVZi = hash.md5("test_CtlNPj")
    log.inf("hash:", h_eLOVZi)
    pt(m_ZQoBcE)
    

    arr_Jzqrcy = [69, 19+1, 66+2]
    map_KxYWeu = {"a": arr_Jzqrcy[0], "b": arr_Jzqrcy[1]}
    pt(map_KxYWeu["a"])
    

    cl Base_oWrGEx {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_AuHXVX <- Base_oWrGEx {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_OLxVYM = Child_AuHXVX(43, 97)
    pt(obj_OLxVYM.get_id())
    

    cl Base_YkxrZq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_uIokDF <- Base_YkxrZq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_pMReQm = Child_uIokDF(24, 14)
    pt(obj_pMReQm.get_id())
    

    fn calc_ozQzOV(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_AnuqTA = calc_ozQzOV(77, 21+1)
    pt(res_AnuqTA)
    
