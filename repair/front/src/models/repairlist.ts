import request from '~/global/request'

export default class Repairlist {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/repairlist',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/repairlist',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/repairlist',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/repairlist',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/repairlist/${id}`
        })

        return res
    }

    static async search(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/repairlist/search',
            params: params
        })

        return res
    }    
}
