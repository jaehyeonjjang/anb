import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
import { getCurrentInstance } from "vue"

import axios from 'axios'
import moment from 'moment'

import { Category, Repair } from "~/models"

export default class util {
    static login(store: any, res: any) {
        store.commit('setLogin', { token: res.token, user: res.user })
    }

    static isNull(value: string) {
        try {
            if (value == null || value == undefined || value == '') {
                return true
            } else {
                return false
            }
        } catch (e) {
            return true
        }
    }

    static getExt(value: string) {
        return value.split('.').pop()
    }

    static forceUpdate() {
        const instance = getCurrentInstance()
        instance?.proxy?.$forceUpdate()
    }

    static money(num: number) {
        num = this.getInt(num)
        var regexp = /\B(?=(\d{3})+(?!\d))/g
        return num.toString().replace(regexp, ',')
    }

    static moneyfloat(num: number) {
        num = parseFloat(num)
        num = Math.floor(num * 100) / 100
        var regexp = /\B(?=(\d{3})+(?!\d))/g
        return num.toString().replace(regexp, ',')
    }

    static moneyempty(num: number) {
        num = this.getInt(num)

        if (num == 0) {
            return ''
        }

        var regexp = /\B(?=(\d{3})+(?!\d))/g
        return num.toString().replace(regexp, ',')
    }

    static makeDate(year: number, month: number, day: number) {
        return this.pad(year, 4) + '-' + this.pad(month, 2) + '-' + this.pad(day, 2);
    }

    static area(num: number) {
        num = Math.floor(num * 1000) / 1000
        var regexp = /\B(?=(\d{3})+(?!\d))/g
        return num.toString().replace(regexp, ',')
    }

    static clone(obj: any) {
        return JSON.parse(JSON.stringify(obj))
    }

    static replaceAll(str: string, searchStr: string, replaceStr: string) {
        return str.split(searchStr).join(replaceStr);
    }

    static nl2br(str: string) {
        if (this.isNull(str)) {
            return ''
        }
        return this.replaceAll(str, '\n', '<BR/>')
    }

    static getDate(d: any) {
        if (this.isNull(d) || d == '') {
            return ''
        }

        d.setTime(d.getTime() + 9 * 60 * 60 * 1000);
        return moment(d).format('YYYY-MM-DD')
    }

    static getYear(d: string) {
        if (this.isNull(d)) {
            return ''
        }

        let temp = d.split('-')
        return temp[0]
    }

    static info(str: string) {
        ElMessage({ message: str, duration: 1000 })
    }

    static error(str: string) {
        ElMessage.error({ message: str, duration: 1000 })
    }

    static confirm(title: string, successFunc: any, cancelFunc: any) {
        ElMessageBox.confirm(
            title,
            '',
            {
                dangerouslyUseHTMLString: true,
                confirmButtonText: '확인',
                cancelButtonText: '취소',
                type: 'warning',
            }
        ).then(() => {
            if (successFunc != undefined && successFunc != null) {
                successFunc()
            }
        }).catch(() => {
            if (cancelFunc != undefined && cancelFunc != null) {
                cancelFunc()
            }
        })
    }

    static messagebox(title: string, content: string, func: any) {
        ElMessageBox.alert(content, title, {
            confirmButtonText: '확인',
            callback: (action: Action) => {
                if (func != undefined && func != null) {
                    func()
                }
            },
        })
    }

    static calculatePrice(direct: number, labor: number, cost: number) {
        direct = this.getFloat(direct)
        labor = this.getFloat(labor)
        cost = this.getFloat(cost)

        let k = labor * 7.9 / 100
        let l = (direct + labor) * 5.5 / 100
        let m = labor * 3.75 / 100
        let n = labor * 0.79 / 100
        let o = labor * 3.23 / 100
        let p = labor * 4.5 / 100

        let q = o * 8.51 / 100
        let r = labor * 2.3 / 100
        let s = (direct + labor) * 3.09 / 100
        let t = (direct + labor + cost) * 0.3 / 100
        let u = (direct + labor + cost) * 0.07 / 100
        let v = (direct + labor + cost) * 6 / 100
        let w = (labor + cost + v) * 15 / 100

        let x = direct + labor + cost + k + l + m + n + o + p + q + r + s + t + u + v + w
        let y = x * 10 / 100

        return parseInt(x + y)
    }

    static calculatePriceRate(direct: number, labor: number, cost: number, rate: number, parcelrate: number) {
        direct = this.getFloat(direct)
        labor = this.getFloat(labor)
        cost = this.getFloat(cost)
        rate = this.getFloat(rate)
        parcelrate = this.getFloat(parcelrate)

        let k = labor * 7.9 / 100
        let l = (direct + labor) * 5.5 / 100
        let m = labor * 3.75 / 100
        let n = labor * 0.79 / 100
        let o = labor * 3.23 / 100
        let p = labor * 4.5 / 100

        let q = o * 8.51 / 100
        let r = labor * 2.3 / 100
        let s = (direct + labor) * 3.09 / 100
        let t = (direct + labor + cost) * 0.3 / 100
        let u = (direct + labor + cost) * 0.07 / 100
        let v = (direct + labor + cost) * 6 / 100
        let w = (labor + cost + v) * 15 / 100

        let x = direct + labor + cost + k + l + m + n + o + p + q + r + s + t + u + v + w
        let y = x * 10 / 100

        let ret = (x + y)

        if (rate != 0.0 && rate != 100.0) {
            ret *= rate / 100.0
        }

        if (parcelrate != 0.0 && parcelrate != 100.0) {
            ret *= parcelrate / 100.0
        }

        return parseInt(ret)
    }

    static calculateRepair(direct: number, labor: number, cost: number, rate: number, parcelrate: number, count: number, percent: number) {
        direct = this.getFloat(direct)
        labor = this.getFloat(labor)
        cost = this.getFloat(cost)
        rate = this.getFloat(rate)
        parcelrate = this.getFloat(parcelrate)
        count = this.getFloat(count)
        percent = this.getFloat(percent)

        /*
        let k = labor * 7.9/100
        let l = (direct + labor) * 5.5/100
        let m = labor * 3.75/100
        let n = labor * 0.79/100
        let o = labor * 3.23/100
        let p = labor * 4.5/100

        let q = o * 8.51/100
        let r = labor * 2.3/100
        let s = (direct + labor) * 3.09/100
        let t = (direct + labor + cost) * 0.3/100
        let u = (direct + labor + cost) * 0.07/100
        let v = (direct + labor + cost) * 6/100
        let w = (labor + cost + v) * 15/100

        let x = direct + labor + cost + k + l + m + n + o + p + q + r + s + t + u + v + w
        let y = x * 10/100

        let ret = (x+y)

        if (rate != 0.0 && rate != 100.0) {
            ret *= rate / 100.0
        }

        if (parcelrate != 0.0 && parcelrate != 100.0) {
            ret *= parcelrate / 100.0
        }
        */

        let ret = this.calculatePriceRate(direct, labor, cost, rate, parcelrate)

        ret *= count
        ret *= percent / 100.0

        return parseInt(ret)
    }

    static getInt(value: any) {
        if (value == undefined || value == null) {
            return 0
        }

        if (typeof value == 'string') {
            value = value.replace(/,/g, '')
        }

        /*
        if (typeof value == 'number') {
            return value
        }
        */

        const ret = parseInt(value)
        if (isNaN(ret)) {
            return 0
        }

        return ret
    }

    static getFloat(value: any) {
        if (value == undefined || value == null) {
            return 0
        }

        if (typeof value == 'string') {
            value = value.replace(/,/g, '')
        }

        if (typeof value == 'number') {
            if (value == 0) {
                return 0.0
            }

            if (String(value).indexOf('.') == -1) {
                value = `${value}.0`
            } else {
                return value
            }
        }

        const ret = parseFloat(value)
        if (isNaN(ret)) {
            return 0
        }

        return ret
    }

    static getFloatFixed(value: any) {
        if (value == undefined || value == null) {
            return 0
        }

        if (typeof value == 'string') {
            value = value.replace(/,/g, '')
        }

        if (typeof value == 'number') {
            if (value == 0) {
                return 0.0
            }

            if (String(value).indexOf('.') == -1) {
                value = `${value}.0`
            } else {
                return Number(value.toFixed(12))
            }
        }

        const ret = parseFloat(value)
        if (isNaN(ret)) {
            return 0.0
        }

        return Number(ret.toFixed(12))
    }

    static fixed(value: number, size: number) {
        let mul = parseInt(Math.pow(10, size))
        let num = Math.floor(value * mul) / mul

        return num
    }

    _loading = null

    static loading(value: boolean) {
        if (value == true) {
            this._loading = ElLoading.service({
                lock: true,
                text: 'Loading',
                background: 'rgba(0, 0, 0, 0.7)',
            })
        } else {
            this._loading.close()
        }
    }

    static pad(value: any, width: number) {
        let str = '' + value;
        return str.length >= width ? str : new Array(width - str.length + 1).join('0') + str;
    }

    static async getCategoryTree(apt: number, title: string) {
        let res = await Category.find({
            apt: apt,
            orderby: 'c_order,c_id'
        })

        if (res.items == null) {
            res.items = []
        }

        let items = res.items

        let categorys = [{ label: title, value: 0 }]

        for (let i = 0; i < items.length; i++) {
            let item = items[i]

            if (item.level == 1) {
                categorys.push({
                    label: item.name,
                    value: item.id,
                    children: []
                })

                let index = categorys.length - 1

                for (let j = 0; j < items.length; j++) {
                    let item2 = items[j]

                    if (item2.parent == item.id) {
                        categorys[index].children.push({
                            label: item2.name,
                            value: item2.id,
                            children: []
                        })

                        let index2 = categorys[index].children.length - 1

                        for (let j = 0; j < items.length; j++) {
                            let item3 = items[j]

                            if (item3.parent == item2.id) {
                                categorys[index].children[index2].children.push({
                                    label: item3.name,
                                    value: item3.id,
                                    children: []
                                })
                            }
                        }

                    }
                }
            }
        }

        return { allcategorys: items, categorys: categorys }
    }

    static async getRepair(id: number) {
        const res = await Repair.get(id)
        return res.item
    }

    static getCategoryChildren(items: any) {
        if (items == null) {
            return null
        }

        let arr = [items]
        if (items.children != null) {
            items.children.forEach((item: any) => arr.push(...this.getCategoryChildren(item)))
        }

        return arr
    }

    static findCategory(id: number, items: any) {
        if (items == null) {
            return null
        }

        let ret = items.find((item: any) => item.value == id)

        if (ret != null) {
            return ret
        }

        for (let i = 0; i < items.length; i++) {
            if (items[i].children == null) {
                continue;
            }

            ret = this.findCategory(id, items[i].children)
            if (ret != null) {
                break
            }
        }

        return ret
    }

    static getPlanyears(value: number) {
        if (value == 0) {
            return 49
        }

        return value - 1
    }

    static download(store: any, url: string, filename: string) {
        axios.get(import.meta.env.VITE_REPORT_URL + url, {
            responseType: 'blob',
            headers: {
                Authorization: 'Bearer ' + store.state.token
            }
        }).then(response => {
            const url = window.URL.createObjectURL(new Blob([response.data]))
            const link = document.createElement('a')
            link.href = url
            link.setAttribute('download', filename);
            document.body.appendChild(link)
            link.click()
        }).catch(exception => {
            alert("파일 다운로드 실패");
        });
    }

    static convertDate(value: any) {
        if (this.isNull(value) || value == '') {
            return ''
        }

        if (typeof value == 'string' || typeof value == 'number') {
            let d = new Date(value)
            return util.getDate(d)
        }

        if (!this.isNull(value) && value != 0) {
            return util.getDate(value)
        }

        return ''
    }

    static convertDBDate(value: any) {
        if (this.isNull(value) || value == '') {
            return '1000-01-01'
        }

        if (typeof value == 'string' || typeof value == 'number') {
            let d = new Date(value)
            return util.getDate(d)
        }

        if (!this.isNull(value) && value != 0) {
            return util.getDate(value)
        }

        return '1000-01-01'
    }

    static viewDate(value: string) {
        if (this.isNull(value) || value == '' || value == '0000-00-00' || value == '1000-01-01') {
            return ''
        }

        let temp = value.split('-')

        return `${temp[0]}년 ${temp[1]}월 ${temp[2]}일`
    }

    static reversPrice(directInt: number) {
        let value = util.getFloat(directInt)

        value = value * 10.0 / 11.0
        let direct = value * 100.0 / (100.0 + 5.5 + 3.09 + 0.3 + 0.07 + 6.0 + 0.9)

        return util.getInt(direct)
    }

    static getImagePath(filename: string, dir: string) {
        if (dir != undefined && dir != '') {
            filename = `${dir}/${filename}`
        }

        return import.meta.env.VITE_REPORT_URL + '/webdata/' + filename
    }

    static getCurrentDate() {
        let d = new Date()
        return this.makeDate(d.getFullYear(), d.getMonth() + 1, d.getDate())
    }
}
