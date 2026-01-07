import request from '~/global/request'

export default class Periodicdata {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicdata',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicdata',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicdata',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicdata',
            params: params
        })

        console.log(res)

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicdata/${id}`
        })

        return res
    }
}
