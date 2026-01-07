import request from '~/global/request'

export default class Oldimage {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/oldimage',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/oldimage',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/oldimage',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/oldimage',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/oldimage/${id}`
        })

        return res
    }
}
