import request from '~/global/request'

export default class User {
    static levels = ['권한', '관리자', '총관리자', '사이트 관리자', '사이트 총관리자']
    static aptlevels = ['권한', '관리자', '총관리자'] 
    static levelTypes = ['', 'info', 'success', 'warning', 'danger']

    static getLevelType(value: number) {
        return this.levelTypes[value]
    }

    static getLevel(value: number) {
        return this.levels[value]
    }

    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/user',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/user',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/user',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/user',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/user/${id}`
        })

        return res
    }

    static async countByLoginid(id: string) {
        const res = await request({
            method: 'GET',
            url: `/api/user/count/loginid/${id}`
        })

        return res.count
    }    
}
