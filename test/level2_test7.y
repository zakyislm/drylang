
    asn fn async_task_dutOLu(x) {
        rev x * 2
    }
    uni async_task_dutOLu(66)
    
    asn fn worker_DdZxoS(y) {
        pt("working on", y)
    }
    mul 2 worker_DdZxoS(75)
    
    awt
    

    cl Base_CKuapI {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_PKNGRo <- Base_CKuapI {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CvDNRY = Child_PKNGRo(57, 73)
    pt(obj_CvDNRY.get_id())
    

    cl Base_svAObB {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_cMnsDN <- Base_svAObB {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_tFFRYB = Child_cMnsDN(21, 90)
    pt(obj_tFFRYB.get_id())
    

    val_aEMihB = 38 % 3
    on (val_aEMihB) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_aEMihB)
    
