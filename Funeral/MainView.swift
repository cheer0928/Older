struct MainView: View {
    var body: some View {
        VStack {
            // 顶部导航栏
            HStack {
                Button(action: { /* 返回 */ }) {
                    Image(systemName: "xmark")
                }
                Spacer()
                Text("安息之选")
                    .font(.title)
                    .fontWeight(.medium)
                Spacer()
                Button(action: { /* 更多选项 */ }) {
                    Image(systemName: "ellipsis")
                }
            }
            .padding()
            
            // 主标题区域
            ZStack {
                RoundedRectangle(cornerRadius: 12)
                    .fill(Color(hex: "#FFF4E6"))
                
                HStack {
                    VStack(alignment: .leading) {
                        Text("尊严告别")
                            .font(.system(size: 28))
                            .fontWeight(.bold)
                            .foregroundColor(Color(hex: "#5D4037"))
                        
                        Text("专业贴心服务")
                            .font(.subheadline)
                            .foregroundColor(Color(hex: "#8D6E63"))
                    }
                    Spacer()
                    
                    // 温馨图标
                    Image("peaceful_symbol")
                        .resizable()
                        .frame(width: 80, height: 80)
                }
                .padding()
            }
            .frame(height: 120)
            .padding(.horizontal)
            
            // 服务选项区域
            // ... 其他UI元素
        }
    }
} 