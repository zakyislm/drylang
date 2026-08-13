
    val_RpShCC = 85 % 3
    on (val_RpShCC) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_RpShCC)
    

    fn calc_fexuGP(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_ULvpVz = calc_fexuGP(77, 7+1)
    pt(res_ULvpVz)
    

    sum_mgmDDB = 0
    lp 6 {
        sum_mgmDDB = sum_mgmDDB + 1
        if (sum_mgmDDB > 100) { done }
    }
    pt(sum_mgmDDB)
    

    cl Base_IBwTiD {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_JaoTmw <- Base_IBwTiD {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_TeLefF = Child_JaoTmw(28, 64)
    pt(obj_TeLefF.get_id())
    

    sum_WwmZsP = 0
    lp 9 {
        sum_WwmZsP = sum_WwmZsP + 1
        if (sum_WwmZsP > 100) { done }
    }
    pt(sum_WwmZsP)
    
