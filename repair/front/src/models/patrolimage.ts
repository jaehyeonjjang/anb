import request from '~/global/request'

export default class Patrolimage {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/patrolimage',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/patrolimage',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/patrolimage',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/patrolimage',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/patrolimage/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/patrolimage/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/patrolimage/find/apt/${apt}`
        })

        return res
    }

    static async countByPatrol(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/patrolimage/count/patrol/${apt}`
        })

        return res.count
    }

    static async findByPatrol(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/patrolimage/find/patrol/${apt}`
        })

        return res
    }
}
