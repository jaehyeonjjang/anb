import request from '~/global/request'

export default class Periodicimage {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicimage',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicimage',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicimage',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicimage/batch',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicimage',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicimage/${id}`
        })

        return res
    }

    static async process(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicimage/process',
            data: item
        })

        return res
    }    
}
