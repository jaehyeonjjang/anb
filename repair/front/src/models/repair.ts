import request from '~/global/request'

export default class Repair {
    static types = ['구분', '재수립', '검토조정', '타업체작업']
    static typeTypes = ['', 'success', 'warning', 'info']

    static getTypeType(value: number) {
        return this.typeTypes[value]
    }

    static getType(value: number) {
        return this.types[value]
    }

    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/repair',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/repair',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/repair',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/repair/batch',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/repair',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/repair/${id}`
        })

        return res
    }

    static async lastdate(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/repair/lastdate/${id}`
        })

        return res
    }

    static async updateStatusById(status: number, id: number) {
        let item = {
            status,
            id
        }

        const res = await request({
            method: 'PUT',
            url: '/api/repair/statusbyid',
            data: item
        })

        return res
    }

    static async change(id: number) {
        const item = {
            id: id
        }
        const res = await request({
            method: 'POST',
            url: '/api/repair/change',
            data: item
        })

        return res
    }
}
