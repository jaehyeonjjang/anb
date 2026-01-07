import request from '~/global/request'

export default class Reviewbasic {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/reviewbasic',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/reviewbasic',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/reviewbasic',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/reviewbasic',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/reviewbasic/${id}`
        })

        return res
    }            
}
