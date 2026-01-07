import request from '~/global/request'

export default class Upload {    
    static async excel(apt: number, params: any) {
        const res = await request({
            method: 'GET',
            url: `/api/upload/excel/${apt}`,
            params: params
        })

        return res
    }

    static async assistance(params: any) {
        const res = await request({
            method: 'POST',
            url: `/api/upload/assistance`,
            data: params
        })

        return res
    }

    static async diff(repair: number, params: any) {
        const res = await request({
            method: 'GET',
            url: `/api/upload/diff/${repair}`,
            params: params
        })

        return res
    }

    static async diffupdate(params: any) {
        const res = await request({
            method: 'POST',
            url: `/api/upload/diffupdate`,
            data: params
        })

        return res
    }
}
