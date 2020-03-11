package xyz.hoper.user.client

import org.springframework.stereotype.Component
import java.util.concurrent.TimeUnit
import io.grpc.ManagedChannel
import io.grpc.ManagedChannelBuilder

@Component
class GrpcClientMananer {
    fun getChannel(host: String?, port: Int): io.grpc.ManagedChannel? {
        return io.grpc.ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .disableRetry()
                .idleTimeout(5, TimeUnit.SECONDS)
                .build()
    }
}