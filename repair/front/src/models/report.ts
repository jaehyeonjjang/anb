import request from '~/global/request'

export default class Report {    
    static async total(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/report/total/${apt}`
        })

        return res
    }

    static async summary(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/report/summary/${apt}`
        })

        return res
    }

    static async plan(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/report/plan/${apt}`
        })

        return res
    }
}
