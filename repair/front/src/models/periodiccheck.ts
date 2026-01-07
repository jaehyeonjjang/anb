import request from '~/global/request'

export default class Periodiccheck {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodiccheck',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodiccheck',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodiccheck',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodiccheck',
            params: params
        })

        if (res.items == null) {
            res.items = []
        }

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodiccheck/${id}`
        })

        return res
    }
}
