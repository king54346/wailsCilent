
const data = [
    {
        path: '/game',
        name: 'Game',
        meta: { title: '游戏', icon: 'Management' }
    },
    {
        path: '/order',
        name: 'Order',
        meta: { title: '工单', icon: 'Management' },
        children: [
            {
                path: 'manage',
                name: 'OrderManage',
                meta: { title: '工单管理', icon: 'Tools' }
            },
        ]
    },
    {
        path: '/annon',
        name: 'Annon',
        meta: { title: '公告', icon: 'Management' },
        children: [
            {
                path: 'manage',
                name: 'AnnonManage',
                meta: { title: 'Annon管理', icon: 'Tools' }
            },
        ]
    },
    // {
    //     path: '/mysql',
    //     name: 'MySQL',
    //     meta: { title: 'MySQL', icon: 'List' },
    //     children: [
    //         {
    //             path: 'sqlAnalysis',
    //             name: 'sqlAnalysis',
    //             meta: { title: 'SQL分析', icon: 'Tools' }
    //         },
    //         {
    //             path: 'sqlSlow',
    //             name: 'sqlSlow',
    //             meta: { title: 'SQL慢日志', icon: 'Histogram' }
    //         },
    //     ]
    // },
    //

    //
    // {
    //     path: '/nginx',
    //     name: 'Nginx',
    //     meta: { title: 'Nginx', icon: 'TrendCharts' },
    //     children: [
    //         {
    //             path: 'nginx',
    //             name: 'nginx',
    //             meta: { title: 'nginx分析', icon: 'TrendCharts' }
    //         },
    //     ]
    // },

]



export default data;